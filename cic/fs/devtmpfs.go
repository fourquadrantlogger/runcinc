package fs

var devtmpfs = MountConfig{
	Target: "/dev",
	Fstype: "devtmpfs",
	Source: "devtmpfs",
	Options: []string{
		"nosuid",
		"strictatime",
		"mode=755",
		"size=65536k",
	},
}
