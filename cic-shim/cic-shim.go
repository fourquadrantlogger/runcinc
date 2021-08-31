package cic_shim

import (
	"context"
	"runcic/cic"
)

type RuncicShim struct {
	cic.Runcic
	cancel context.CancelFunc
}
