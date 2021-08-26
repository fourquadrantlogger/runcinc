package cic

import (
	"math/rand"
	"runcic/utils"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}
func newName() string {
	return utils.GetRandomString(3+rand.Intn(5)) + "_" + utils.GetRandomString(3+rand.Intn(3))
}
func newID() string {
	return utils.GetRandomString2(12)
}
