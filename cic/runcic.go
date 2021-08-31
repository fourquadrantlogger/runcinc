package cic

import (
	"context"
	"os"
	"runcic/containerimage/common"
	"strings"
	"time"
)

type Runcic struct {
	ParentRootfs    *os.File
	cancel          context.CancelFunc
	Name            string
	CicVolume       string
	ContainerID     string
	Images          []*common.Image
	Command         []string
	Envs            []string
	Started         time.Time
	ImagePullPolicy ImagePullPolicy
}

func (r *Runcic) ImageArray() (imgs []string) {
	for i := 0; i < len(r.Images); i++ {
		imgs = append(imgs, r.Images[i].Image)
	}
	return
}
func (r *Runcic) Roorfs() (path string) {
	path = OverlayRoot + string(os.PathSeparator) + r.Name
	return
}

func (r *Runcic) mountops() string {
	lower := make([]string, 0)
	for i := 0; i < len(r.Images); i++ {
		lower = append(lower, r.Images[i].Lower...)
	}
	mountops := strings.Join([]string{
		"lowerdir=" + strings.Join(lower, ":"),
		"upperdir=" + r.CicVolume + string(os.PathSeparator) + "up",
		"workdir=" + r.CicVolume + string(os.PathSeparator) + "work",
	}, ",")
	return mountops
}
