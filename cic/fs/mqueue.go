package fs

var mqueue = MountConfig{
	Target: "/dev/mqueue",
	Fstype: "mqueue",
	Source: "mqueue",
	Options: []string{
		"nosuid",
		"noexec",
		"nodev"},
}
