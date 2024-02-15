package Go

import (
	"fmt"
	"strconv"
)

// https://www.codewars.com/kata/5b077ebdaf15be5c7f000077/train/go

func IfYouCantSleepJustCountSheepMySelf(num int) string {
	var result string
	for i := 1; i <= num; i++ {
		result += strconv.Itoa(i) + " sheep..."
	}
	return result
}

func IfYouCantSleepJustCountSheep(num int) string {
	var result string
	for i := 1; i <= num; i++ {
		result += fmt.Sprintf("%v sheep...", i)
	}
	return result
}
