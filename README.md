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

… TODO 



## 安装

### go-micro

[官方Github](https://github.com/asim/go-micro)

go-micro的repo其实有些混乱，以及除了go-micro之外，还有一个micro的repo，两者之间的关系我暂时还没搞懂（似乎micro可安装后对go-micro进行查看）

安装go-micro的v3版本，其实之前有段时间好像叫nitro来着，但最后正式版还是回归了go-micro

```
go get -u github.com/asim/go-micro/v3
```

如果出问题记得设置

```
go env -w GO111MODULE=on
```

go-micro的v1和v2都是`github.com/micro/go-micro/[v2]`，不知道为什么反而v3变回了个人的repo路径

#### go-micro plugins

go-micro在v1和v2采用`go-plugins`的一个单独repo [Github](https://github.com/microhq/go-plugins)，但在v3的时候又重新回到了`go-micro/plugins`下面[Github](https://github.com/asim/go-micro/tree/master/plugins)

没有成功获得插件正确安装方法，只会傻傻得一个个安装。。

v2的安装命令，注意替换需要的插件的路径

```
go get -u github.com/micro/go-plugins/registry/consul/v2
```

v3

```
go get -u github.com/asim/go-micro/plugins/registry/consul/v3
```

### protobuf

#### 安装protobuf

protobuf介绍自行搜索，类似json但占用更小的空间

在[官方Github](https://github.com/protocolbuffers/protobuf)里的release下载最新的包，选择对应的os，比如我下载的是`protoc-3.14.0-win64.zip`（不懂为什么一群人教程下32位）

放到想要的安装目录下，比如我的`C:\Program Files\Protobuf`，把bin目录加入环境变量

命令行

```
>protoc --version
libprotoc 3.14.0
```

#### 安装protoc-gen-go

protobuf默认是不支持go的编译的，因此需要安装`protoc-gen-go`来支持编译为go文件

```
go get -u github.com/golang/protobuf/protoc-gen-go
```

#### 安装protoc-gen-micro

protoc-gen-micro是用于生成针对micro可用的文件 [Github](https://github.com/asim/go-micro/tree/master/cmd/protoc-gen-micro)

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



