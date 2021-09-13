package cmd

import (
	"strings"
	"testing"
)

func TestExecute(t *testing.T) {
	var args = strings.Split(`--copyenv  --name myedi --cap-add CAP_SYS_ADMIN --cap-drop CAP_DDD  --env vara=a  --cicvolume /data/edi/ codingcorp-docker.pkg.coding.net/cloud-studio-next/docker/workspace-golang:2021.14.2  bash`, " ")
	var args_ = []string{}
	for i := 0; i < len(args); i++ {
		if args[i] != "" {
			args_ = append(args_, args[i])
		}
	}
	parse(args_)
}
