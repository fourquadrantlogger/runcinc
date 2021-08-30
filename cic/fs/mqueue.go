package fs

import "syscall"

var mqueue = MountConfig{
	Target: "/dev/mqueue",
	Fstype: "mqueue",
	Source: "mqueue",
	Flags:  syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV,
}
