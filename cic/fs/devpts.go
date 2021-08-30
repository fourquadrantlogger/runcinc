package fs

var devpts = MountConfig{
	Target: "/dev/pts",
	Fstype: "devpts",
	Source: "devpts",
	Options: []string{
		"noexec=0",
		"newinstance",
		"ptmxmode=0666",
		"mode=0620",
		"gid=5"},
}

var ptmx = MountConfig{
	Target:  "/dev/ptmx",
	Source:  "/dev/pts/ptmx",
	Options: []string{"bind"},
}
