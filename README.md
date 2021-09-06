#runcic

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
