package fs

var devtmpfs = MountConfig{
	Target: "/dev",
	Fstype: "tmpfs",
	Source: "tmpfs",
	Options: []string{
		"mode=755",
		"size=65536k",
	},
}

var shm = MountConfig{
	Target: "/dev/shm",
	Fstype: "tmpfs",
	Source: "shm",
	Options: []string{
		"nodev",
		"mode=1777",
		"size=65536k"},
}
