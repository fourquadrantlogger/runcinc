name=$1
cicroot=/runcic/overlay/${name}
runcic runin --name=$name ${@:3}
umount ${cicroot}/dev/mqueue
umount ${cicroot}/dev/shm
umount ${cicroot}/dev/pts
umount ${cicroot}/dev
umount ${cicroot}/proc
umount ${cicroot}/sys
umount ${cicroot}