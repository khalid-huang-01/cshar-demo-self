## 介绍
​        游戏进程和GSE通过grpc通信，通信pb的协议见gameserver_grpcsdk_service.proto和gse_grpcsdk_service.proto。其中gameserver_grpcsdk_service.proto定义了三个服务接口，需要由游戏进程来实现，gse_grpcsdk_service.proto定义的服务接口GSE来实现，游戏进程需要在合适的时机调用对应的接口。



## 游戏进程集成流程
### 1. 实现gameserver_grpcsdk_service.proto定义的三个接口
​	游戏进程在启动后，需要实现grpc server，并实现gameserver_grpcsdk_service.proto定义的三个接口。
#### 1.1 OnHealthCheck
​	GSE会每隔一分钟调用一次该接口，用来获取游戏进程的健康状态，游戏进程需要通过该接口返回当前进程的健康状态。 返回结构如下：
 
```
message HealthCheckResponse {
    bool healthStatus = 1;
}
```

 	如果进程健康则healthStatus为true，否则为false

#### 1.2 OnStartGameServerSession

​	当游戏开发者通过调用腾讯云CreateGameServerSession等相关接口生成GameServerSession后，GSE会通过该接口把GameServerSession回传给游戏进程

​	游戏进程收到请求后，需要保存该GameServerSession对应的信息，GameServerSession对应的结构见pb

#### 1.3 OnProcessTerminate

​	在缩容阶段或者健康检查持续失败，GSE会调用该接口，告诉游戏进程需要结束进程，游戏进程收到请求后，需要结束其上在1.2步骤承载的游戏会话，并调用GSE实现的两个接口TerminateGameServerSession（见2.7）和ProcessEnding（见2.8），告诉GSE需要结束该游戏会话和结束进程



### 2. 在合适时机调用GSE相关接口
#### 2.1 ProcessReady（必须调用）

​	在游戏进程启动后，准备就绪可以承载游戏进程后，需要告诉调用该接口告诉GSE进程已就绪，请求的参数如下：

```
message ProcessReadyRequest {
    repeated string logPathsToUpload = 1;
    int32 clientPort = 2;
    int32 grpcPort = 3;
}
```

​	其中，logPathsToUpload为游戏进程需要上传的日志路径，GSE会将其指定的日志路径文件上传到腾讯云，供游戏开发者下载；clientPort为游戏客户端要链接的端口；grpcPort为在1中实现gameserver_grpcsdk_service.proto 定义的服务指定的端口，该端口供GSE来链接

#### 2.2 ActivateGameServerSession（必须调用）

​	1.2中收到GSE的回调后，需要调用该接口告诉GSE来激活对应的GameServerSession。请求的参数如下：

```
message ActivateGameServerSessionRequest{
    string gameServerSessionId = 1;
    int32 maxPlayers = 2;
}
```

​	其中gameServerSessionId为1.2回调中的GameServerSession对象的里的一个字段，用了唯一标记一个游戏会话，maxPlayers为该游戏会话最大允许加入的玩家数

#### 2.3 AcceptPlayerSession （必须调用）

​	当玩家加入游戏后，游戏进程需要调用该接口来告诉GSE某个玩家已经加入。请求参数如下：

```
message AcceptPlayerSessionRequest {
    string gameServerSessionId = 1;
    string playerSessionId = 2;
}
```

​	GSE通过该接口传入的gameServerSessionId和playerSessionId参数来校验该玩家的合法性。其中playerSessionId为游戏开发者通过调用腾讯云JoinGameServerSession返回的玩家在对应的游戏会话中的唯一id。

#### 2.4 RemovePlayerSession （必须调用）

​	当玩家退出游戏后，游戏进程需要调用该接口来告诉GSE某个玩家已经退出。请求参数如下：

```
message RemovePlayerSessionRequest {
    string gameServerSessionId = 1;
    string playerSessionId = 2;
}
```

​	GSE收到该请求后，会更新对应的游戏会话的当前玩家数，以允许其他玩家加入进来

#### 2.5 DescribePlayerSessions （根据业务可选）

​	游戏进程可以通过该接口，获取当前已经加入到其上承载的游戏会话的玩家信息，请求参数如下：

```
message DescribePlayerSessionsRequest {
   string gameServerSessionId = 1;
   string playerId = 2;
   string playerSessionId = 3;
   string playerSessionStatusFilter = 4;
   string nextToken = 5;
   int32 limit = 6 ;
}
```

返回结果见 DescribePlayerSessionsResponse，具体见pb定义

#### 2.6 UpdatePlayerSessionCreationPolicy （根据业务可选）

​	游戏进程可以通过该接口，更新当前游戏会话玩家加入的策略，请求参数如下：

```
message UpdatePlayerSessionCreationPolicyRequest {
   string gameServerSessionId = 1;
   string newPlayerSessionCreationPolicy = 2;
}
```

​	其中newPlayerSessionCreationPolicy为要更新后的策略，可选值有**ACCEPT_ALL**（接受所有新玩家会话） 和**DENY_ALL** （拒绝所有新玩家会话）

#### 2.7 TerminateGameServerSession （必须调用）

​	游戏进程需要调用该接口通知GSE其上承载的游戏会话已结束，请求参数如下：

```
message TerminateGameServerSessionRequest {
    string gameServerSessionId = 1;
}
```

#### 2.8 ProcessEnding （必须调用）

​	游戏进程需要调用该接口通知GSE该游戏进程正在关闭，请求参数如下：

```
message ProcessEndingRequest {
}
```

#### 2.9 ReportCustomData （根据业务可选）

​	游戏进程可以调用该接口告知GSE游戏进程的自定义数据，请求参数如下：

```
message ReportCustomDataRequest {
    int32 currentCustomCount = 1 ;
    int32 maxCustomCount = 2;
}
```



### 3 其他

在游戏进程通过grpc调用2里面的相关接口到GSE时，需要在grpc请求的meta里添加两个字段

| 字段      | 含义                                      | 类型   |
| --------- | ----------------------------------------- | ------ |
| pid       | 当前游戏进程的pid                         | string |
| requestId | 当前请求的requestId，用了唯一标记一次请求 | string |

