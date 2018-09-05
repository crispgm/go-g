package sanitizer

import "regexp"

// RemoveDelimiters 过滤空格、标点、控制字符、Unicode各类标点、数学符号
func RemoveDelimiters(word string) string {
	rules := regexp.MustCompile(`[[:space:][:punct:][:cntrl:]\pP\pS]+`)
	return rules.ReplaceAllLiteralString(word, "")
}
