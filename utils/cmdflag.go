package utils

import (
	"os"
	"runcic/cic"
)

//ParseCmdFlag runcic run 命令，前部分参数传递给runcic，后部分参数全部传给子容器image，因此需要单独实现一个parse
func ParseCmdFlag() (cfg cic.CicConfig) {

	for i := 0; i < len(os.Args); i++ {

	}
	return
}
