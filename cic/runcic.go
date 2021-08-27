package cic

import (
	"os"
	"runcic/containerimage/common"
	"strings"
	"time"
)

type Runcic struct {
	Name            string
	CicVolume       string
	ContainerID     string
	Image           *common.Image
	Command         []string
	Envs            []string
	Started         time.Time
	ImagePullPolicy ImagePullPolicy
}

func (r *Runcic) Roorfs() (path string) {
	path = OverlayRoot + string(os.PathSeparator) + r.Name
	return
}

func (r *Runcic) mountops() string {
	mountops := strings.Join([]string{
		"lowerdir=" + strings.Join(r.Image.Lower, ":"),
		"upperdir=" + r.CicVolume + string(os.PathSeparator) + "up",
		"workdir=" + r.CicVolume + string(os.PathSeparator) + "work",
	}, ",")
	return mountops
}
