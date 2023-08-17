/*
   @Author Ted
   @Since 2023/8/5 19:43
*/

package common

import (
	"douyinProject/config"
	"douyinProject/log"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/redigo"
	"strconv"

	//"github.com/garyburd/redigo/redis"    //另一种redis客户端，由不同团队开发的
	"github.com/gomodule/redigo/redis"
	"time"
)

var (
	redisClient *redis.Pool
	rsy         *redsync.Redsync //分布式锁库，redsync.Redsync 可以使用 redisClient 来与Redis服务器通信并执行锁定操作
	//锁过期时间
	lockExpiry = 2 * time.Second
	//获取锁失败重试时间间隔
	retryDelay = 500 * time.Millisecond
	//值过期时间
	valueExpire  = 86400 //一天
	ErrMissCache = errors.New("miss Cache")
	//锁设置，
	option = []redsync.Option{
		redsync.WithExpiry(lockExpiry),     //设置锁过期时间
		redsync.WithRetryDelay(retryDelay), //重试时间间隔
	}
)

func RedisInit() {
	conf := config.GetConfig()
	network := conf.Redis.NetWork
	address := conf.Redis.Host
	port := conf.Redis.Port
	auth := conf.Redis.Auth
	host := fmt.Sprintf("%s:%s", address, port)
	redisClient = &redis.Pool{
		MaxIdle:     10,                //允许最大的空闲连接数，当空闲连接数超过该值时，多余的连接会被关闭
		MaxActive:   0,                 //表示连接池中允许的最大活跃（非空闲）连接数，如果设置为 0，则表示没有限制
		IdleTimeout: 240 * time.Second, //超过该时间的空闲连接将被关闭
		Wait:        true,              //为 true，表示当连接池中没有可用连接时，其他请求会等待，直到有连接可用

		//使用 redis.Dial 函数来创建 Redis 连接，并指定相关的网络、主机、密码和数据库。
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(network, host,
				redis.DialPassword(auth),
				redis.DialDatabase(0), //连接哪个数据库
			)
			if err != nil {
				log.Info("redis连接失败,", err)
				return nil, err
			}
			return c, err
		},
	}
	redisClient.Get().Do("flushdb")     //FLUSHDB 命令，用于清空当前选择的数据库中的所有数据
	sync := redigo.NewPool(redisClient) //创建了一个 Redis 连接池，并将其与 redisClient 关联起来
	rsy = redsync.New(sync)             //创建了一个分布式锁、管理器，使用之前创建的 Redis 连接池 sync
	//log.Println(redisClient.Dial == nil)
	log.Info("redis连接成功")

}

func Exists(key string) bool {
	conn := redisClient.Get() //从客户端获取一个连接
	defer conn.Close()

	//conn.Do的返回值是interface{}，可以用redis.Bool把它转成布尔；成功就是bool，
	_, err := conn.Do("EXISTS", key)
	if err != nil {
		return false
	}

	return true
}

// redis的set
func CacheSet(key string, value any) error {
	conn := redisClient.Get()
	defer conn.Close()
	data, err := json.Marshal(value) //序列化成二进制存储
	if err != nil {
		return err
	}
	_, err = conn.Do("set", key, data, "ex", valueExpire) //1.不区分大小写；  2.是ex不是nx..
	if err != nil {
		return err
	}
	return nil
}

func CacheGet(key string) ([]byte, error) {
	conn := redisClient.Get()
	defer conn.Close()
	obj, err := redis.Bytes(conn.Do("get", key))
	if err != nil {
		return nil, err
	} else if len(obj) == 0 { //不存在该key
		return nil, ErrMissCache
	}
	return obj, nil
}

//	func CacheHSet(key, mkey, value string) error {
//		conn := redisClient.Get()
//		defer conn.Close()
//		//data, err := json.Marshal(value) //序列化成二进制存储
//		_, err := conn.Do("hset", key, mkey, value)
//
//		//for _, val := range value { //第二个参数才是value，第一个是index
//		//
//		//	if err != nil {
//		//		return err
//		//	}
//		//
//		//	if err != nil {
//		//		return err
//		//	}
//		//}
//		return err
//	}
func CacheHSet(key, mkey string, value ...interface{}) error {
	conn := redisClient.Get()
	defer conn.Close()

	for i, d := range value {
		data, err := json.Marshal(d)
		if err != nil {
			return nil
		}

		_, err = conn.Do("HSET", key, mkey+"_"+strconv.Itoa(i), data)
		if err != nil {
			return err
		}
	}
	return nil
}
func CacheHGet(key, mkey string) ([]byte, error) {
	conn := redisClient.Get()
	defer conn.Close()

	data, err := redis.Bytes(conn.Do("HGET", key, mkey))

	//fmt.Printf("data:%v", data)
	if err != nil {
		return []byte{}, err
	}
	if len(data) == 0 {
		return []byte{}, ErrMissCache
	}
	return data, nil
}

func CacheHGetAll(key string) (map[string]string, error) {
	conn := redisClient.Get()
	defer conn.Close()
	result, err := redis.Values(conn.Do("hgetall", key)) //转成byte切片的切片
	fmt.Printf("result的类型为:%T", result)
	var mp map[string]string
	for i := 0; i < len(result); i += 2 {
		mp[result[i].(string)] = result[i+1].(string)
	}
	log.Info("获取的map为：", mp)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return mp, nil
}

func CacheLPush(key string, data ...interface{}) error {
	conn := redisClient.Get()
	defer conn.Close()
	for _, binaryTmp := range data {
		value, _ := json.Marshal(binaryTmp)
		_, err := conn.Do("lpush", key, value)
		if err != nil {
			log.Error("lpush出错,", err.Error())
			return err
		}
	}
	return nil
}

func CacheLGetAll(key string) ([][]byte, error) {
	conn := redisClient.Get()
	defer conn.Close()
	data, err := redis.ByteSlices(conn.Do("lrange", key, "0", "-1")) //转成byte切片的切片
	if err != nil {
		log.Error(err.Error())
		return [][]byte{}, err
	}
	return data, nil
}

func CacheLRem(key string, data any) error {
	conn := redisClient.Get()
	defer conn.Close()
	_, err := conn.Do("LREM", key, "0", data)
	return err
}

func GetLock(key string) (*redsync.Mutex, error) {
	mutex := rsy.NewMutex(key+"_lock", option...)
	if err := mutex.Lock(); err != nil {
		return mutex, err
	}
	return mutex, nil
}

func UnLock(mutex *redsync.Mutex) error {
	if _, err := mutex.Unlock(); err != nil {
		return err
	}
	return nil
}
