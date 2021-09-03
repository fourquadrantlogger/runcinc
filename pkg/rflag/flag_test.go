package rflag

import (
	"fmt"
	"testing"
)

func TestFindIndex(t *testing.T) {
}

func TestParseFlag(t *testing.T) {
	fmt.Println(ParseFlag([]string{
		"--env", "fw=af",
		"--cicvolume", "/cic/w",
		"--cicimage", "/image",
		"--envcopy",
		"codingcorp-docker.pkg.coding.net/cloud-studio-next/docker/editor-server-image:2021.141,codingcorp-docker.pkg.coding.net/cloud-studio-next/docker/editor-server-image:2021.141",
		"bash", "-c", "sleep", "10h"},
		func([]string) int { return 0 },
		func([]string) int { return -1 },
		[]string{"env"},
	))
}
