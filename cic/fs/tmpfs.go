package fs

var shm = MountConfig{
	Target: "/dev/shm",
	Fstype: "tmpfs",
	Source: "shm",
	Options: []string{
		"noexec",
		"nodev",
		"mode=1777",
		"size=65536k"},
}
