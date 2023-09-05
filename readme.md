# 抖音基础版（字节跳动青训项目）
## 一、项目介绍
- 本抖音项目是基于grpc通讯工具开发的高性能微服务，不仅使用gin作为业务层框架，gorm框架作为持久层框架，还使用预编译sql防止sql注入，同时该项目结合连接池技术来构建连接工厂和复用grpc连接来提高系统的性能，这样可以有效的处理高并发场景下的挑战，还可以通过减少频繁创建和销毁grpc连接带来的性能开销
- 项目服务地址:https://1024code.com/codecubes/jpyi9rm 
- 项目地址:https://github.com/fineCoderWithLove/douyin-base

## 二、项目实现
 ###  2.1技术选型
- gin:提供grpc服务使用protobuf进行数据传输。
- JWT:token生成和权限的校验
- Gorm:对Mysql之心ORM操作，Go-redis：操作Redis对频繁更改的数据进行缓存以便更快的响应。
- Redis:对点赞/取消赞，视频的喜欢量/评论量，用户的喜欢量，总点赞量缓存Redis中，设置定时任务，并且使数据同步到数据库中。
- Zap:高性能日志打印
- ffmpeg:进行视频取帧，作为视频的封面
- 七牛云:使用七牛云做对象存储，用来存储视频，图片等静态资源。
- pprof:使用pprof进行性能测试

  ### 2.2架构设计
  由于项目的耦合度不高，所以采用微服务架构来缓解服务器的压力，项目分为api层，业务服务层，数据层
- api层负责鉴权和分发请求调用远程服务来返回数据
- 业务层负责与数据库进行交互和逻辑处理

![image](https://github.com/fineCoderWithLove/douyin-base/assets/122780660/aaee18fe-bfb1-4216-9bdb-a3033c74c116)



  ## 三、测试结果
  ### 3.1功能测试
  ### 3.2性能测试
  1. 我们使用pprof进行性能监测，因为每次请求grpc都会产生连接和销毁连接造成服务的性能消耗，思考后我把grpc的连接设置成一个全局变量，后来发现这个全局变量有一个问题，在并发情况下，用同一个全局变量会导致读写错误。
  2. 经过思考，我设置了互斥锁的全局变量，但是性能提升不是很明显。
  3. 经过搜索引擎查询资料，最后利用线程池技术，简单工厂设计模式设计出了一个GrpcFactory工厂，每次只需要调用工厂就可以返回连接配合利用grpc的keep-alive使得grpc的连接开销变小。性能测试图如下:
#### 优化前
![image](https://github.com/fineCoderWithLove/douyin-base/assets/122780660/1adacf86-77dd-4a8b-90b2-69f0fa645102)

#### 优化后
![image](https://github.com/fineCoderWithLove/douyin-base/assets/122780660/746d799e-bbeb-457c-9fd7-6f558e4ef482)

## 四、项目总结与反思
### 4.1目前存在的问题
1. 上传视频会上传到本地的内路径中
2. 聊天记录存储到mysql中可能导致查询聊天记录mysql压力过大
### 4.2已经识别的优化项
1. 判断user和video是否存在的时候，可以直接从redis中判断增加速度
2. 应该将聊天记录缓存到redis中{token:create_time}的形式，因为前端需要不断获取到最晚消息的发布时间
### 4.3架构演进的可能性
1. 分库分表
2. 分布式
3. 配置更多微服务组件
### 4.4项目中的反思和总结
- 代码应该尽可能优雅的写法，让以后的改动是方便的，应该满足开放封闭原则。
- 一个优秀的程序员应该让别的程序员更好的工作，我的队友给我提供了很多的工具，让我工作更加高效。
- 测试是一个项目的重点，没有测试的软件是不合格的，而测试的关键则是边界点的问题。
- 每一个同步的位置都是并发情况下容易发生错误的地方，都要加上互斥锁。
- 一个项目应该敢为极致，在自己力所能及的地方做到最好，应该尝试多种可能性，寻找最好的解决办法！

## 五、部署
1. 部署需要ffmpeg环境
2. 改变每个模块中global的mysql连接和redis连接
3. Linux下执行命令./run.sh
