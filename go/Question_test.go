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

func TestCenturyFromYear(t *testing.T) {
	test1 := CenturyFromYearMySelf(1990)
	assert.Equal(t, 20, test1)
	test2 := CenturyFromYearMySelf(1705)
	assert.Equal(t, 18, test2)
	test3 := CenturyFromYearMySelf(631933)
	assert.Equal(t, 6320, test3)
	test4 := CenturyFromYearMySelf(93)
	assert.Equal(t, 1, test4)
}

func TestIfYouCantSleepJustCountSheep(t *testing.T) {
	test1 := IfYouCantSleepJustCountSheep(2)
	assert.Equal(t, "1 sheep...2 sheep...", test1)
	test2 := IfYouCantSleepJustCountSheep(0)
	assert.Equal(t, "", test2)
	test4 := IfYouCantSleepJustCountSheep(1)
	assert.Equal(t, "1 sheep...", test4)
}

func TestTotalAmountOfPoints(t *testing.T) {
	test1 := TotalAmountOfPoints([]string{"1:0","2:0","3:0","4:0","2:1","3:1","4:1","3:2","4:2","4:3"})
	assert.Equal(t, 30, test1)
	test2 := TotalAmountOfPoints([]string{"1:1","2:2","3:3","4:4","2:2","3:3","4:4","3:3","4:4","4:4"})
	assert.Equal(t, 10, test2)
	test4 := TotalAmountOfPoints([]string{"0:1","0:2","0:3","0:4","1:2","1:3","1:4","2:3","2:4","3:4"})
	assert.Equal(t, 0, test4)
}

func TestHighAndLow(t *testing.T) {
	test1 := HighAndLow("8 3 -5 42 -1 0 0 -9 4 7 4 -4")
	assert.Equal(t, "42 -9", test1)
	test2 := HighAndLow("1 2 3")
	assert.Equal(t, "3 1", test2)
}

func TestReturningStrings(t *testing.T) {
	test1 := ReturningStrings("Ryan")
	assert.Equal(t, "Hello, Ryan how are you doing today?", test1)
}
