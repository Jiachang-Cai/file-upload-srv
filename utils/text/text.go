package text

import "strings"

// 替换文本为自定义字符
func ReplaceText(text, char string, start, end uint64) string {
	l := uint64(len(text))
	if end < start {
		return text
	}
	if start > l || end > l {
		return text
	}
	charLen := end - start
	charList := make([]string, charLen+1)
	return text[0:start] + strings.Join(charList, char) + text[end:]
}
