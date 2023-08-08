
# TikTok
字节青训营的抖音项目

## 更新日志
8.1: 添加视频点赞，喜欢列表功能（但是刷新时红心会消失，这个的解决办法不太清楚；尝试修改is_favorite字段但还是没用....各位可以帮忙看一看）\
8.2: 添加"评论操作","评论列表"功能 \
8.8: 添加了redis功能，响应的配置在config.yaml里，根据自己的情况修改

## 简要说明
**目前完成所有基础接口：视频流、用户注册、用户登录、用户信息、投稿接口、发布列表。**

部分错误处理可能不是太完善，大致功能测试过，主要功能没问题；不过测试样例比较少，可能某些特殊情况会出现bug的情况。日志、错误处理还不太完善。

其中的配置信息config.yaml需要依据各自的情况进行修改，包括mysql数据库、Minio分布式存储系统； 以及service包下的videoservice的GetImage的ffmpeg的路径。（Minio主要用于视频的云端存储，生成的url存储在本地数据库）


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

文件说明:

main.go:主程序入口，包含init初始化配置

routes.go: 统一处理各自请求


### 部分注意事项
1. 本地测试时上传视频请求时间较长，展示视频的请求时间也比较长（应该是请求Minio比较花时间，敬请耐心等待）
2. Minio建议使用windows下的（尝试过linux和windows下，踩了很多坑，后者好用一点）
3. 数据库表不需要提前创建，只需要创建数据库就行；表在运行DBinit方法时会自动创建（gorm的映射）
4. Minio、ffmpeg需要先配置准备好，具体可以查看网上教程
5. 其中可能会有一些坑，如果遇到问题，可以互相交流一下
6. 测试时可以用postman模拟发请求
