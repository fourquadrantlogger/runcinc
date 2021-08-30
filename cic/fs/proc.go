package fs

var proc = MountConfig{
	Source: "proc",
	Target: "/proc",
	Fstype: "proc",
	Options: []string{
		"nosuid",
		"noexec",
		"nodev",
	},
}
