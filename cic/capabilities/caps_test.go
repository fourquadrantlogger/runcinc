package capabilities

import (
	"syscall"
	"testing"
)

func TestNew(t *testing.T) {
	cs := []string{
		"CAP_SYS_ADMIN",
		"CAP_CHOWN",
		"CAP_UNKNOWN",
		"CAP_UNKNOWN2"}
	conf := Capabilities{
		Bounding:    cs,
		Effective:   cs,
		Inheritable: cs,
		Permitted:   cs,
		Ambient:     cs,
	}
	caps, _ := New(&conf)

	if len(caps.caps) != len(capTypes) {
		t.Errorf("expected %d capability types, got %d: %v", len(capTypes), len(caps.caps), caps.caps)
	}

	//for _, cType := range capTypes {
	//	if i := len(caps.caps[cType]); i != 1 {
	//		t.Errorf("expected 1 capability for %s, got %d: %v", cType, i, caps.caps[cType])
	//		continue
	//	}
	//}
	err := caps.ApplyCaps()
	if err != nil {
		t.Errorf("ApplyCaps error, got %+v: ", err.Error())
	}
	err = syscall.Mount("/data/image", "/data/cicimg", "bind", syscall.MS_BIND|syscall.MS_REC, "")
	if err != nil {
		t.Errorf("Mount error, got %+v: ", err.Error())
	}
}
