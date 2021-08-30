package fs

import "syscall"

var sys = MountConfig{
	Target:  "/sys",
	Fstype:  "sysfs",
	Source:  "sysfs",
	Flags:   syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV | syscall.MS_RDONLY,
	Options: []string{},
}
