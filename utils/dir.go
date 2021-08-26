package utils

import "os"

func Mkdirp(dirpath string) {
	os.MkdirAll(dirpath, 0775)
}
