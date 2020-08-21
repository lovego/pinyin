package pinyin

import (
	"regexp"
	"strings"
)

var (
	filterRe = regexp.MustCompile("[^a-zA-Z0-9\u4e00-\u9fa5]")
)

// 获取所有汉字的首字母组合， 比如"长城ABC"，返回"CCABC"
// -keepSpecial 保留特殊符号
func Initials(s string, keepSpecial bool) string {
	if !keepSpecial {
		s = filterRe.ReplaceAllString(s, "")
	}
	for i := 0; i < len(dict); i += 2 {
		s = strings.Replace(s, dict[i], dictInitials(dict[i+1]), -1)
	}
	return strings.ToUpper(s)
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
