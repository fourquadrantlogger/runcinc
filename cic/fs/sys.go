package fs

var sys = MountConfig{
	Target: "/sys",
	Fstype: "sysfs",
	Source: "sysfs",
	Options: []string{
		"nosuid",
		"noexec",
		"nodev",
		"ro"},
}
