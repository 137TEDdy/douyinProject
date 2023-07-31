
# TikTok
字节青训营的抖音项目

## 简要说明
    目前完成所有基础接口：视频流、用户注册、用户登录、用户信息、投稿接口、发布列表。
    部分错误处理可能不是太完善，大致功能测试过，主要功能没问题；不过测试样例比较少，可能某些特殊情况会出现bug的情况
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