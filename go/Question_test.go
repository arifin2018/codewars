package Go

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCSVRepresentationOfArray(t *testing.T) {
	test1 := CSV_representation_of_array([][]int{
		{0, 1, 2, 3, 45},
		{10, 11, 12, 13, 14},
		{20, 21, 22, 23, 24},
		{30, 31, 32, 33, 34}},
	)
	// fmt.Println(test1)
	assert.Equal(t, "0,1,2,3,45\n10,11,12,13,14\n20,21,22,23,24\n30,31,32,33,34", test1)

}

func TestMultipleOfIndex(t *testing.T) {
	test1 := multipleOfIndex([]int{
		22, -6, 32, 82, 9, 25,
	},
	)
	// fmt.Println(test1)
	assert.Equal(t, []int{-6, 32, 25}, test1)
}

func TestYouCantCodeUnderPressure_1(t *testing.T) {
	test1 := YouCantCodeUnderPressure_1(1)
	assert.Equal(t, 2, test1)
}

func TestCheckForFactor(t *testing.T) {
	test1 := CheckForFactor(10, 2)
	assert.Equal(t, true, test1)
}

func TestConvertAStringToAnArray(t *testing.T) {
	test1 := ConvertAStringToAnArray("Robin Singh")
	assert.Equal(t, []string{"Robin", "Singh"}, test1)
}

func TestAbbreviateATwoWordName(t *testing.T) {
	test1 := AbbreviateATwoWordName("sam harris")
	assert.Equal(t, "S.H", test1)
}

func TestIsItAPalindrome(t *testing.T) {
	test1 := IsItAPalindrome("Abba")
	assert.Equal(t, true, test1)
	test2 := IsItAPalindrome("hello")
	assert.Equal(t, true, test2)
}
