package containerimage

import (
	"os/exec"
	"runcic/containerimage/common"
	"runcic/containerimage/docker"
	"runcic/containerimage/podman"
)

var defaultImageDriver common.ImageDriver
var drivemap = map[string]common.ImageDriver{
	"podman": &podman.Podman{
		Root: "/var/lib/containers/storage",
	},
	"docker": &docker.Docker{},
}

func init() {
	for cmd, drive := range drivemap {
		_, err := exec.LookPath(cmd)
		if err == nil {
			defaultImageDriver = drive
		}
	}
}
func Driver() common.ImageDriver {
	return defaultImageDriver
}
func SetDriver(driver common.ImageDriver) {
	defaultImageDriver = driver
}
