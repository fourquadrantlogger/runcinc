name=$1
cicroot=/runcic/overlay/${name}
runcic runin --name=$name ${@:3}
umount ${cicroot}/dev/mqueue
umount ${cicroot}/dev/shm
umount ${cicroot}/dev/pts
umount ${cicroot}/dev
umount ${cicroot}/proc
umount ${cicroot}/sys
set -e
umount ${cicroot}/.PlnPyKFp4CRfFtgC1_run
umount ${cicroot}
rm -rf ${cicroot}