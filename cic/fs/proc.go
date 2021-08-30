package fs

import "syscall"

var proc = MountConfig{
	Source: "proc",
	Target: "/proc",
	Fstype: "proc",
	Flags:  syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV,
}
