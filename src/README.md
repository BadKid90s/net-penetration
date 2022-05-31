# 内外穿透

## 说明

内网穿透是我们在进行网络连接时的一种术语，也叫做NAT穿透，

## 原理

1. 本机与隧道建立TCP连接 
2. 隧道与公网服务建立TCP连接
3. 公网发请求进行转换通过隧道把请求转发到内外
4. 内网响应数据到隧道，把数据转发到公网

## 实现

### 服务端
1. 创建控制中心监听
2. 创建用户请求监听
3. 创建隧道监听
4. 监听隧道消息转发到用户请求

### 客户端

1. 连接控制中心
2. 连接隧道
3. 连接本地服务
4. 进行消息转发
5. 转发本地服务消息到隧道

## 流程

```
客户端              本地服务             隧道                   控制中心            用户请求          
3.启动客户端监听     4.启动本地服务        2.启动隧道监听        0.启动控制中心监听     1.启动用户请求监听
                                                5.服务端实现用户请求和隧道的数据转发
                                           <-------------------------------------------->
           6.客户端连接服务端
--------------------------------------------------------------->
           7.服务端与就客户端连接完成，实现心跳，保持连接
<---------------------------------------------------------------

 8.客户端连接本地服务
---------------------->
 9.客户端连接隧道服务
---------------------------------------->
                   10.客户端实现本地服务与隧道的数据转发
                     <------------------>
                                          11.用户公网访问，转发给隧道
                                        <-------------------------------------------------
                     12.客户端监听到隧道数据，把隧道转发给本地服务                 
                    <-------------------    
                     12.客户端监听到本地服务数据，把本地服务数据转发给隧道                                     
                    ------------------->
                                       12.服务端监听到隧道数据，把隧道数据转发给用户公网请求                                     
                                       -------------------------------------------------->                                 
                                          
```



