package cmd

import (
	cic_shim "runcic/cic-shim"
)

func cmdWait() {

	cic_shim.Wait(cfg)
}
