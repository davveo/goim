### goim即时通讯系统

#### 设计架构图

<img src="https://cdn.jsdelivr.net/gh/davveo/imagehousing/img/image-20210226134818807.png" alt="image-20210226134818807" style="zoom:50%;" />

#### 分层

##### api

- 对外提供http接口：用户登陆注册、鉴权等信息

##### logic

- 接收api消息并发送到mq

##### connect

- websocket链接管理
- 消息发送与接收

##### mq

##### consumer

#### 技术点

- 用户发送消息，如何找到接受者消息链接？

  > 

- ws链接如何保活？

- 如何实现注册与自动发现？

#### 参考项目

- https://github.com/LockGit/gochat
- https://github.com/woodylan/go-websocket
- https://github.com/link1st/gowebsocket







































