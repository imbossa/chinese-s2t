package chinese_s2t

import (
	lang "github.com/imbossa/chinese-s2t/lang"
	"strings"
	"unicode"
)

// TString 繁体字
var TString = lang.Chinese["T"]

// SString 简体字
var SString = lang.Chinese["S"]

func T2s(t string) string {
	var isChinese bool
	var result = ""

	for _, code := range t {
		isChinese = unicode.Is(unicode.Scripts["Han"], code)
		if !isChinese {
			result += string(code)
			continue
		}

		index := strings.Index(TString, string(code))
		if index != -1 {
			result += SString[index : index+3]
		} else {
			result += string(code)
		}
	}

	return result
}

func S2t(t string) string {
	var isChinese bool
	var result = ""

	for _, code := range t {
		isChinese = unicode.Is(unicode.Scripts["Han"], code)
		if !isChinese {
			result += string(code)
			continue
		}

		index := strings.Index(SString, string(code))
		if index != -1 {
			result += TString[index : index+3]
		} else {
			result += string(code)
		}
	}

	return result
}
