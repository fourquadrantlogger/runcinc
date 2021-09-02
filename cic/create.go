package cic

import (
	"runcic/utils"
)

func (r *Runcic) mergeEnv(cmdenv []string) {
	for _, img := range r.Images {

		for _, v := range img.Env {

			r.Envs = append(r.Envs, v)
		}
	}
	r.Envs = append(r.Envs, cmdenv...)
	envmap := utils.ParseEnvs(r.Envs)
	r.Envs = make([]string, 0)
	for k, v := range envmap {
		r.Envs = append(r.Envs, k+"="+v)
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
