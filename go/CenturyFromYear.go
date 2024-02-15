package Go

import (
	"strconv"
)

//https://www.codewars.com/kata/5a3fe3dde1ce0e8ed6000097/solutions/go

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
