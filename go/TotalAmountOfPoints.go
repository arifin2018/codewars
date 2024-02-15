package Go

import (
	"strconv"
)

// https://www.codewars.com/kata/5bb904724c47249b10000131/train/go

func TotalAmountOfPointsMySelf(games []string) int {
	var result int
	for _, v := range games {
		ourTeam,_ := strconv.Atoi(string(v[0]))
		opponentsTeam,_ := strconv.Atoi(string(v[2]))
		if ourTeam > opponentsTeam {
			result += 3
		}else if ourTeam == opponentsTeam {
			result += 1
		}
	}
	return result
}

func TotalAmountOfPoints(games []string) int {
	var result int
	for _, v := range games {
		if v[0] > v[2] {
			result += 3
		}else if v[0] == v[2] {
			result += 1
		}
	}
	return result
}