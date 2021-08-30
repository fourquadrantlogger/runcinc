package fs

import "syscall"

var shm = MountConfig{
	Target: "/dev/shm",
	Fstype: "tmpfs",
	Source: "shm",
	Flags:  syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV,
	Options: []string{
		"mode=1777",
		"size=65536k"},
}
