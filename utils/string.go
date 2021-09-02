package utils

func MergeStrings(old, new []string) (merged []string) {
	for _, vValue := range new {
		exist := false
		for _, pathValue := range old {
			if pathValue == vValue {
				exist = true
				break
			}
		}
		if exist {
			//fmt.Println("ignore",old,vValue)
		} else {
			old = append(old, vValue)
			//append
			//fmt.Println("append",old,vValue)
		}
	}
	merged = old
	return
}
