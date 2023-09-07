![image](https://github.com/137TEDdy/douyinProject/assets/120385461/aa446fe7-d99f-4848-b810-28c4c2f4c63f)
# TikTok
字节青训营的抖音项目


## 简要说明
**目前完成所有接口： \
基础接口: 视频流、用户注册、用户登录、用户信息、投稿接口、发布列表。 \
互动接口: 关注操作、关注列表、粉丝列表、好友列表、消息记录、发送消息 \
社交接口: 点赞操作、点赞列表、评论操作、评论列表**  \

其中的配置信息config.yaml需要依据各自的情况进行修改，包括mysql数据库、Minio分布式存储系统； 以及service包下的videoservice的GetImage的ffmpeg的路径。（Minio主要用于视频的云端存储，生成的url存储在本地数据库）

## 技术选型:
**语言**：golang 

**HTTP框架**：gin 

**数据库**：mysql,  redis 

**ORM框架**：gorm   

**OSS服务**：minio  

**视频抽帧**：ffmpeg   

**项目部署**：docker 




## 其它技术的使用
1.viper库配置管理

2.zap日志框架

3.使用mysql事务保证一致性

4.引入协程异步处理

5.使用自增状态码，error处理的更高效

6.密码加密处理、token加密
密码加密使用bcrypt密码哈希函数，将密码与随机生成的盐值进行混合，并重复计算哈希值来增加密码的安全性，可以有效地防止常见的密码攻击，如彩虹表攻击和暴力破解。
token使用HS256算法进行签名




## 包结构与文件说明
1.common:通用包，包括数据库启动，jwt鉴权，封装一些response

2.config: 读取config.yaml的配置信息并封装起来

3.controller: 视图层，处理请求，调用业务层方法等

4.service: 业务逻辑层，调用持久层repo的方法等

5.repo: 持久层，与数据库打交道

6.model: 模型层，数据库信息与go的结构体的映射

7.middleware: 中间件（类似java里的拦截器）

8.minioHandler: 处理minio存储的相关操作，包括初始化，创建桶，上传文件

9.utils: 提供一些工具方法

10.test: 单元测试

文件说明:

main.go:主程序入口，包含init初始化配置

routes.go: 统一处理各自请求


