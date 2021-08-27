package cic

import (
	"runcic/utils"
)

func (r *Runcic) Create() (err error) {
	//merge envs
	for _, v := range r.Image.Env {
		//todo 需要深入研究，为啥envs是字符串
		r.Envs = append(r.Envs, v)
	}
	//merge cmd
	if len(r.Command) == 0 {
		r.Command = r.Image.Cmd
	}

	//create name &id
	r.ContainerID = newID()
	r.Name = newName()

	utils.Mkdirp(OverlayRoot)
	utils.Mkdirp(r.CicVolume)
	return
}
