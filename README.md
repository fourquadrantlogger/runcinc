#runcic

## 问题背景

在k8s环境中，容器在销毁之后，除了volume内的数据其它容器内改动，都会被丢弃
k8s容器经常会发生迁移节点，这也会间接导致容器内非volume数据的丢失

## runcic设计目标

runcic设计用来实现在容器内运行镜像作为子容器
当子容器随着父容器销毁，通过子容器diff通过volume实现持久化，避免了容器销毁后改动丢失

## 为什么runcic支持组合多个image作为一个容器运行？

这个设计相当于将y层的image，切分成x1层+x2层+x3...的多个image，其中x1+x2+x3=y
用户无需重写dockerfile组装出新的image，就可以把多个image合并为一个新的image，作为容器运行
## build

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

## usage

```
sh runcic.sh myruncic runin   \
--copyenv   \
--env vara=a   \
--cicvolume=/data/edi/  \
--authfile=/run/secret/mysecret/.dockerconfigjson  \
golang:latest,redis:latest \ 
go env
```

```
apiVersion: v1
kind: Pod
metadata:
  name: mypod
spec:
  containers:
  - name: mypod
    image: runcic
    cmd:["sh","runcic.sh","myruncic","runin","--copyenv","--authfile=/run/secret/mysecret/.dockerconfigjson","golang:latest,redis:latest","go","env"]
    volumeMounts:
    - name: foo
      mountPath: "/run/secret/mysecret"
      readOnly: true
  volumes:
  - name: foo
    secret:
      secretName: mysecret
```
## args
sh runcic.sh myapp   runcic.sh负责回收runcic创建的mount路径

--copyenv  传递父容器环境变量到子容器 bool类型

--authfile  当以volume挂载kubernetes.io/dockerconfigjson类型的secret时，可以以上述形式传递

支持传递多个image，运行的时候会把image的lowerdir按顺序拼接起来，越前列的layer优先级越高


#### runcic的overlay文件系统
runcic的overlay文件系统，按image顺序，联合了多个image的lowerdir

例如，runcic run ubuntu:latest,centos:latest

ubuntu的lowerdir为l1:l2

centos的lowerdir为l3:l4

则，按照ubuntu:latest,centos:latest的顺序，组成的runcic overlay文件系统为
l1:l2:l3:l4

按照overlay的规则，https://www.kernel.org/doc/html/latest/filesystems/overlayfs.html

l1是最上层的lowerdir，l4则为最下层的lowerdir
