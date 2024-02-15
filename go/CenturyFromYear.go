package Go

import (
	"strconv"
)

func CenturyFromYearMySelf(year int) int {
	yearIntToString := strconv.Itoa(year)
	result, _ := strconv.Atoi(yearIntToString[:len(yearIntToString)-2])
	if year%100 == 0 {
		return result
	} else {
		return result + 1
	}
}
func CenturyFromYear(year int) int {
	if year%100 == 0 {
		return year / 100
	}
	return year/100 + 1
}
