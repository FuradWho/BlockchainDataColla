# BlockchainDataColla [![License](https://img.shields.io/:license-apache-blue.svg)](https://opensource.org/licenses/Apache-2.0) [![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/go-micro.dev/v4?tab=doc)

`Hyperledger Fabric` 实验性项目

## 文档目录

- [项目背景](#项目背景)
- [安装](#安装)
- [使用](#使用)
- [Badge](#Badge)
- [相关项目](#相关项目)
- [主要项目负责人](#主要项目负责人)
- [参与贡献方式](#参与贡献方式)
- [开源协议](#开源协议)

## 项目背景

基于 `Fabric` 以及`go-micro`设计的消息协同系统。

随着市场产业的发展，企业的规模不断增大，企业内部的各个部门业务范围不断增加，以及企业间的业务需求愈加复杂。这种情况导致了企业间信息沟通孤立、通信变差、消息不同步等问题。为探求一种基于企业之间去中心化协作消息服务的共享模式，试图来解决目前在企业间消息协同过程之中存在的难一致性、难相互信任等问题。通过深入了解分析目前企业间传统的“点对点”协同模式，发现了利用区块链这一技术的本质特征可以有效解决，从而得到了一种基于区块链的去中心化协作消息服务。

本基于区块链的去中心化协作消息服务，使用来自 Linux Foundation 的开源项目，一个模块化区块链框架“Hyperledger Fabric”作为整个区块链的底层网络。利用其要求许可权的网络特点，使得已参与的企业在整个网络之中建立分散式信任，而不是在匿名参与者的网络之中，并对于彼此之间的消息数据进行保密等。让企业用户更加透明、可信的在网络内部进行消息数据的传递、应答、追踪等双向操作。当用户上传消息给指定的接收方时，服务会主动推送相关通知，同时将消息数据上传到区块链之中以保障数据的不可篡改性，一致性。

本协同消息服务同时使用了Go-Micro微服务框架，利用其基于插件化RPC模块设计，实现了区块链服务、CA认证服务、企业服务、以及客户端等模块，从而将整个复杂的服务系统拆分成多个更小，更独立自洽的服务，服务与服务间通过松耦合的形式交互。
同时整个服务提供了整个区块链网络信息的可视化平台，让用户可以去查看区块链网络的状态、已参与网络的企业信息、消息数据历史信息等。使得用户更加直观地了解整个协同系统的变化。


## 安装

### go-micro

[官方Github](https://github.com/asim/go-micro)

`go-micro`的`repo`其实有些混乱，以及除了`go-micro`之外，还有一个`micro`的`repo`

```
go get -u github.com/asim/go-micro/v3
```

```
go env -w GO111MODULE=on
```

`go-micro`的`v1`和`v2`都是`github.com/micro/go-micro/[v2]`

#### go-micro plugins

`go-micro`在`v1`和`v2`采用`go-plugins`的一个单独 `repo` [Github](https://github.com/microhq/go-plugins)，但在v3的时候又重新回到了`go-micro/plugins`下面[Github](https://github.com/asim/go-micro/tree/master/plugins)

`v2`的安装命令，注意替换需要的插件的路径

```
go get -u github.com/micro/go-plugins/registry/consul/v2
```

`v3`

```
go get -u github.com/asim/go-micro/plugins/registry/consul/v3
```

### protobuf

#### 安装 protobuf

`protobuf`介绍自行搜索，类似json但占用更小的空间

在[官方Github](https://github.com/protocolbuffers/protobuf)里的release下载最新的包，选择对应的os，比如我下载的是`protoc-3.14.0-win64.zip`，把bin目录加入环境变量。

命令行：

```
>protoc --version
libprotoc 3.14.0
```

#### 安装protoc-gen-go

`protobuf`默认是不支持`go`的编译的，因此需要安装`protoc-gen-go`来支持编译为`go`文件

```
go get -u github.com/golang/protobuf/protoc-gen-go
```

#### 安装protoc-gen-micro

`protoc-gen-micro`是用于生成针对`micro`可用的文件 [Github](https://github.com/asim/go-micro/tree/master/cmd/protoc-gen-micro)

```
go get -u github.com/asim/go-micro/cmd/protoc-gen-micro/v3
```

#### 测试命令

```
 protoc --plugin=protoc-gen-go=D:\proiect2021\GO\bin\protoc-gen-go.exe --plugin=protoc-gen-micro=D:\proiect2021\GO\bin\protoc-gen-micro.exe --micro_out=E:\projects\BlockchainData
Colla\fabricDeploy\proto --go_out=. Test.proto
```



## 使用



## Badge



### 相关项目



## 主要项目负责人



## 参与贡献方式



## 开源协议



