package rflag

import (
	"regexp"
)

const (
	LongflagPattern  = `^--[A-Za-z0-9][A-Za-z0-9_\.\-]*$`
	ShortflagPattern = `^-[A-Za-z0-9][A-Za-z0-9_\.\-]*$`
)

type IndexFlag func(args []string) (idx int)

func stringshas(strs []string, s string) (idx int) {
	for i := 0; i < len(strs); i++ {
		if strs[i] == s {
			return i
		}
	}
	return -1
}

//ParseFlag
//如果flag存在重复的key，appendkey中的key对value进行append，其它key的情况则value直接替换
func ParseFlag(args []string, appendKey []string) (flags map[string][]string, unknownArgs []string) {
	flagargs := args
	for i := 0; i < len(flagargs); i++ {
		kmatch, _ := regexp.MatchString(LongflagPattern, flagargs[i])
		kmatchshort, _ := regexp.MatchString(ShortflagPattern, flagargs[i])
		var k string
		if kmatch {
			k = flagargs[i][2:]
		} else if kmatchshort {
			k = flagargs[i][1:]
		} else {
			unknownArgs = append(unknownArgs, flagargs[i])
			continue
		}

		if i != len(flagargs)-1 {
			vmatchshort, _ := regexp.MatchString(ShortflagPattern, flagargs[i+1])
			vmatch, _ := regexp.MatchString(LongflagPattern, flagargs[i+1])
			if vmatchshort || vmatch {
				flags[k] = []string{}
				//existflag
			} else {
				v := flagargs[i+1]
				i++
				if stringshas(appendKey, k) >= 0 {
					flags[k] = append(flags[k], v)
				} else {
					flags[k] = []string{v}
				}
			}
		} else {
			flags[k] = []string{}
		}
	}
	return
}
