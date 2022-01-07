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



