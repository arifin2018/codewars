package kyu

import (
	"regexp"
	"strings"
)

func ConvertStringToCamelCase(s string) string {
	var result string
	regexSymbol := regexp.MustCompile("-|_")
	newStr := regexSymbol.ReplaceAllString(s, "|")
	splitData := strings.Split(newStr, "|")
	result += splitData[0]
	for i := 1; i < len(splitData); i++ {
		tempDataConvert := strings.ToUpper(string(splitData[i][0])) + splitData[i][1:len(splitData[i])]
		result += tempDataConvert
	}
	return result
}
