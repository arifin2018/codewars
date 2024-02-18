package Go

import "fmt"

// https://www.codewars.com/kata/51f2d1cafc9c0f745c00037d/train/go

// failed

func StringEndsWith(str, ending string) bool {
	if len(ending) > len(str) {
		return false
	}
	fmt.Println(str)
	fmt.Println(len(str))
	fmt.Println(len(ending))
	return str[len(str)-len(ending):] == ending
	// return strings.HasSuffix(str, ending)
}
