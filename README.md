#runcic

```
sh runcic.sh myapp runin \
 --name myapp  \
 --copyenv
 --image codingcorp-docker.pkg.coding.net/cloud-studio-next/docker/editor-server-image:2021.14.1,codingcorp-docker.pkg.coding.net/cloud-studio-next/docker/workspace-golang:2021.14.2  \ 
 --env vara=a,var2=b   \
 --cicvolume=/data/edi/  \
 --cicimage /image  \
 bash 
```