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
		"--envcopy"},
		[]string{"env"},
	))
}
