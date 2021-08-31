package cic

import (
	"runcic/utils"
)

func (r *Runcic) mergeEnv() {
	for _, img := range r.Images {
		//todo 需要深入研究，为啥envs是字符串
		for _, v := range img.Env {
			//todo 需要深入研究，为啥envs是字符串
			r.Envs = append(r.Envs, v)
		}
	}
}
func (r *Runcic) mergeCmd() {
	//merge cmd,use firstimage cmd notnull
	if len(r.Command) == 0 {
		for i := 0; i < len(r.Images); i++ {
			if len(r.Images[i].Cmd) > 0 {
				r.Command = r.Images[i].Cmd
				break
			}
		}
	}
}
func (r *Runcic) Create() (err error) {
	utils.Mkdirp(OverlayRoot)
	utils.Mkdirp(r.CicVolume)
	return
}
