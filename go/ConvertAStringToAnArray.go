package Go

import "strings"

func ConvertAStringToAnArray(str string) []string {
	return strings.Split(str, " ")
}
