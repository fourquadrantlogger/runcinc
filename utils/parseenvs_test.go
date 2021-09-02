package utils

import (
	"fmt"
	"os"
	"testing"
)

func TestParseEnvs(t *testing.T) {
	envarr := os.Environ()
	envarr = append(envarr, "PATH=/bin:/GOPATH/bin:/NODE/bin")
	envs := ParseEnvs(envarr)
	for k, v := range envs {
		fmt.Println(k, v)
	}
}

func TestMergeEnv(t *testing.T) {
	envs := ParseEnvs(os.Environ())
	MergeEnv(envs, []string{
		"PATH=/bin:/GOPATH/bin:/la/bin",
		"GCCGO=cc",
	})
	for k, v := range envs {
		fmt.Println(k, v)
	}
}
