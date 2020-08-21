package pinyin

import (
	"regexp"
	"strings"
	"unicode"
)

var (
	filterRe = regexp.MustCompile("[^a-zA-Z0-9\u4e00-\u9fa5]")
)

// 获取所有汉字的首字母组合， 比如"长城ABC"，返回"CCABC"
// -keepSpecial 保留特殊符号
func Initials(s string, keepSpecial bool) (res string) {
	if !keepSpecial {
		s = filterRe.ReplaceAllString(s, "")
	}
	for _, word := range split(s) {
		res += word.initials()
	}
	return
}

func dictInitials(s string) string {
	var v string
	for _, item := range strings.Split(s, "\t") {
		if item != `` {
			v += strings.ToUpper(item[0:1])
		}
	}
	return v
}

type word struct {
	V       string
	Chinese bool
}

func split(s string) (ws []word) {
	var temp string
	lastType := 0
	for _, v := range s {
		t := unicode.Is(unicode.Han, v)
		if (t && lastType == 1) || (!t && lastType == -1) || (lastType == 0) {
			temp += string(v)
		} else {
			ws = append(ws, word{
				V:       temp,
				Chinese: lastType == 1,
			})
			temp = string(v)
			lastType = 0
		}
		if t {
			lastType = 1
		} else {
			lastType = -1
		}
	}
	ws = append(ws, word{
		V:       temp,
		Chinese: lastType == 1,
	})
	return
}

func (w word) initials() (res string) {
	if !w.Chinese {
		return strings.ToUpper(w.V)
	}
	chinese := []rune(w.V)
	vL := len(chinese)
	for i := 0; i < vL; {
		var j int
		var contain bool
		for ; j <= vL-i-1; j++ {
			if py, ok := dictMap[string(chinese[i:vL-j])]; ok {
				res += dictInitials(py)
				contain = true
				break
			}
		}
		if !contain {
			res += string(chinese[i : vL-j+1])
			i++
		} else {
			i += vL - i - j
		}
	}
	return
}
