package podman

import (
	"fmt"
	"testing"
)

func TestPodman_Spec(t *testing.T) {
	podMan := &Podman{}
	spec := podMan.Spec("ubuntu:latest")
	fmt.Printf("%v", spec)
}

func TestPodman_Pull(t *testing.T) {
	podMan := &Podman{}
	podMan.Pull("redis:latest")
}
