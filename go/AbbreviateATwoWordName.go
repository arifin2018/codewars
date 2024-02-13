package Go

import (
	"strings"
)

func AbbreviateATwoWordNameMySelf(name string) string {
	var result string
	splitString := strings.Split(name, " ")
	for i := 0; i < len(splitString); i++ {
		result += strings.ToUpper(string(splitString[i][0]))
		if i < len(splitString)-1 {
			result += "."
		}
	}
	return result
}

func AbbreviateATwoWordName(name string) string {
	var result []string
	splitString := strings.Split(name, " ")
	for i := 0; i < len(splitString); i++ {
		result = append(result, strings.ToUpper(string(splitString[i][0])))
	}
	return strings.Join(result, ".")
}
