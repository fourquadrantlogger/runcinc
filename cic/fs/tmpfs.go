package fs

var devtmpfs = MountConfig{
	Target: "/dev",
	Fstype: "tmpfs",
	Source: "tmpfs",
	Options: []string{
		"nosuid",
		"strictatime",
		"mode=755",
		"size=65536k",
	},
}

var shm = MountConfig{
	Target: "/dev/shm",
	Fstype: "tmpfs",
	Source: "shm",
	Options: []string{
		"nosuid",
		"noexec",
		"nodev",
		"mode=1777",
		"size=65536k"},
}
