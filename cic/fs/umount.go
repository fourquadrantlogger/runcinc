package fs

import (
	"github.com/moby/sys/mount"
)

func CleanMount(target string) {
	mount.RecursiveUnmount(target)
}
