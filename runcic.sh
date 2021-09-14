name=$1
cicroot=/runcic/overlay/${name}
runcic runin --name=$name ${@:3}
umount -A  --recursive ${cicroot}
