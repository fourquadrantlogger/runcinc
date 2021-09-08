#runcic

## 问题背景

在k8s环境中，容器在销毁之后，除了volume内的数据其它容器内改动，都会被丢弃
k8s容器经常会发生迁移节点，这也会间接导致容器内非volume数据的丢失


## runcic设计目标

runcic设计用来实现在容器内运行镜像作为子容器
当子容器随着父容器销毁，通过子容器diff通过volume实现持久化，避免了容器销毁后改动丢失


## build

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

## usage

```
sh runcic.sh myedi runin   \
--copyenv   \
--env vara=a   \
--cicvolume=/data/edi/  \
golang:latest \ 
go env
```

## args
sh runcic.sh myapp  
--copyenv  传递父容器环境变量到子容器 bool类型
--image 支持传递多个image，运行的时候会把image的lowerdir按顺序拼接起来，越前列的layer优先级越高

## docker usage 
