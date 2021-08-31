package cic

import (
	"runcic/utils"
)

func (r *Runcic) Create() (err error) {
	//merge envs
	for _, img := range r.Images {
		//todo 需要深入研究，为啥envs是字符串
		for _, v := range img.Env {
			//todo 需要深入研究，为啥envs是字符串
			r.Envs = append(r.Envs, v)
		}
	}

	//merge cmd,use firstimage cmd notnull
	if len(r.Command) == 0 {
		for i := 0; i < len(r.Images); i++ {
			if len(r.Images[i].Cmd) > 0 {
				r.Command = r.Images[i].Cmd
				break
			}
		}
	}

	//create name &id
	r.ContainerID = newID()
	r.Name = newName()

	utils.Mkdirp(OverlayRoot)
	utils.Mkdirp(r.CicVolume)
	return
}
