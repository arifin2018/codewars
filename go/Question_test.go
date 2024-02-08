package Go

import (
	"fmt"
	"testing"
)

func TestCSVRepresentationOfArray(t *testing.T) {
	test1 := CSV_representation_of_array1([][]int{
		{0, 1, 2, 3, 45},
		{10, 11, 12, 13, 14},
		{20, 21, 22, 23, 24},
		{30, 31, 32, 33, 34}},
	)
	fmt.Println(test1)
	// assert.Equal(t, "0,1,2,3,45\n10,11,12,13,14\n20,21,22,23,24\n30,31,32,33,34", test1)

}
