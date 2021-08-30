package fs

import (
	"syscall"
)

var devpts = MountConfig{
	Target: "/dev/pts",
	Fstype: "devpts",
	Flags:  syscall.MS_NOEXEC | syscall.MS_NOSUID,
	Source: "devpts",
	Options: []string{
		"newinstance",
		"ptmxmode=0666",
		"mode=0620",
	},
}

var ptmx = LinkConfig{
	Target: "/dev/ptmx",
	Source: "/dev/pts/ptmx",
}
