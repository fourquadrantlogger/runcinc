package docker

import (
	"fmt"

	"testing"
)

func TestPodman_Spec(t *testing.T) {
	docker := &Docker{}
	spec := docker.Spec("ubuntu:latest")
	fmt.Printf("%v", spec)
}

func TestPodman_Pull(t *testing.T) {
	docker := &Docker{}
	docker.Pull("redis:latest", "")
}
