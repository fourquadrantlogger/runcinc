package containerimage

import (
	"runcic/containerimage/common"
	"runcic/containerimage/podman"
)

var defaultImageDriver common.ImageDriver

func init() {
	defaultImageDriver = &podman.Podman{}
}
func Driver() common.ImageDriver {
	return defaultImageDriver
}
func SetDriver(driver common.ImageDriver) {
	defaultImageDriver = driver
}
