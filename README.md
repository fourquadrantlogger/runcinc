#runcic

## build

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

## usage

```
sh runcic.sh myapp runin \
 --copyenv \
 --image codingcorp-docker.pkg.coding.net/cloud-studio-next/docker/editor-server-image:2021.14.1,codingcorp-docker.pkg.coding.net/cloud-studio-next/docker/workspace-golang:2021.14.2  \ 
 --env vara=a,var2=b   \
 --cicvolume=/data/edi/  \
 --cicimage /image  \
 bash 
```

## args
sh runcic.sh myapp  ;myapp必须和 --name myapp相同名，这是由于runcic.sh 的回收脚本需要name，暂时没做复杂的flag解析支持

--copyenv  传递父容器环境变量到子容器 bool类型

--image 支持传递多个image，运行的时候会把image的lowerdir按顺序拼接起来，越前列的layer优先级越高
