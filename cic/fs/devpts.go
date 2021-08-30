package fs

var devpts = MountConfig{
	Target: "/dev/pts",
	Fstype: "devpts",
	Source: "devpts",
	Options: []string{
		"nosuid",
		"noexec",
		"newinstance",
		"ptmxmode=0666",
		"mode=0620",
	},
}

var ptmx = MountConfig{
	Target:  "/dev/ptmx",
	Source:  "/dev/pts/ptmx",
	Options: []string{"bind"},
}
