package utils

import "os"

func Mkdirp(dirpath string) {
	os.MkdirAll(dirpath, 0775)
}

func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true

}
