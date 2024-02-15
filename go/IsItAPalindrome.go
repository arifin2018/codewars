package Go

import "strings"

func IsItAPalindrome(str string) bool {
	var result string
	for i := len(str) - 1; i >= 0; i-- {
		result += string(str[i])
	}
	return strings.EqualFold(result, str)
}
