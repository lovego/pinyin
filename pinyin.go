package utils

import (
	"github.com/mozillazg/go-pinyin"
)

var firstLetters = getArgs()

func getArgs() pinyin.Args {
	a := pinyin.NewArgs()
	a.Style = pinyin.FirstLetter
	return a
}

// 获取所有汉字的首字母组合， 比如"长城ABC"，返回"CCABC"
func Initials(str string) string {
	var runes = []rune(str)
	var bytes = make([]byte, 0, len(runes))
	for _, r := range runes {
		var b byte
		if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' {
			b = byte(r)
		} else if letters := pinyin.SinglePinyin(
			r, firstLetters,
		); len(letters) > 0 && len(letters[0]) > 0 {
			b = letters[0][0]
		}

		switch {
		case b >= 'a' && b <= 'z':
			b = b - ('a' - 'A')
			bytes = append(bytes, b)
		case b >= 'A' && b <= 'Z':
			bytes = append(bytes, b)
		}
	}
	return string(bytes)
}
