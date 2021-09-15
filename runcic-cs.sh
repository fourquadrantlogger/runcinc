name=$1
FILE=/cic/sized.img
export X_IDE_SPACE_SIZE=5

if [ -f "$FILE" ]; then
    echo "$FILE exist"
else
    dd if=/dev/null of=$FILE bs=8M seek=$[128*$X_IDE_SPACE_SIZE]
    mkfs.ext4 -N 3276800 -E discard -F -F $FILE
fi

if [ ! -d /sizedcic ];then
    mkdir /sizedcic
fi

mount $FILE /sizedcic
cicroot=/runcic/overlay/${name}
rm -rf /var/lib/containers/storage/libpod
runcic runin --name=$name --cicvolume=/sizedcic ${@:3}
umount -A  --recursive ${cicroot}
umount /sizedcic