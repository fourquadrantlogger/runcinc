package fs

import "syscall"

var devtmpfs = MountConfig{
	Target: "/dev",
	Fstype: "devtmpfs",
	Source: "devtmpfs",
	Flags:  syscall.MS_NOEXEC | syscall.MS_STRICTATIME,
	Options: []string{
		"mode=755",
	},
}
