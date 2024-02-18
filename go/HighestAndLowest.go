package Go

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// https://www.codewars.com/kata/554b4ac871d6813a03000035/train/go

func HighAndLow1(in string) string {
	// 2272ms 
	data := strings.Split(in, " ")
	high,_ := strconv.Atoi(string(data[0]))
	low,_ := strconv.Atoi(string(data[0]))
	for _, v := range data {
		tempInt,_ := strconv.Atoi(string(v))
		if high < tempInt {
			high = tempInt
		}
		if low > tempInt {
			low = tempInt
		}
	}
	return fmt.Sprintf("%d %d", high, low)
}

func HighAndLow(in string) string {
	// 2101ms 
	data := strings.Split(in, " ")
	high,_ := strconv.Atoi(string(data[0]))
	low,_ := strconv.Atoi(string(data[0]))
	for i := 0; i < int(math.Ceil(float64(len(data)/2))); i++ {
		tempIntFront,_ := strconv.Atoi(string(data[i]))
		tempIntBack,_ := strconv.Atoi(string(data[len(data)-(i+1)]))
		if high < tempIntFront {
			high = tempIntFront
		}
		if low > tempIntFront{
			low = tempIntFront
		}
		if high < tempIntBack {
			high = tempIntBack
		}
		if low > tempIntBack{
			low = tempIntBack
		}
	}
	return fmt.Sprintf("%d %d", high, low)

}