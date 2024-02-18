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
	test1 := TotalAmountOfPoints([]string{"1:0", "2:0", "3:0", "4:0", "2:1", "3:1", "4:1", "3:2", "4:2", "4:3"})
	assert.Equal(t, 30, test1)
	test2 := TotalAmountOfPoints([]string{"1:1", "2:2", "3:3", "4:4", "2:2", "3:3", "4:4", "3:3", "4:4", "4:4"})
	assert.Equal(t, 10, test2)
	test4 := TotalAmountOfPoints([]string{"0:1", "0:2", "0:3", "0:4", "1:2", "1:3", "1:4", "2:3", "2:4", "3:4"})
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
func TestIsHeGonnaSurvive(t *testing.T) {
	test1 := IsHeGonnaSurvive(10, 5)
	assert.Equal(t, true, test1)
	test2 := IsHeGonnaSurvive(7, 4)
	assert.Equal(t, false, test2)
	test3 := IsHeGonnaSurvive(4, 5)
	assert.Equal(t, false, test3)
	test4 := IsHeGonnaSurvive(100, 40)
	assert.Equal(t, true, test4)
}

func TestStringEndsWith(t *testing.T) {
	test1 := StringEndsWith("", "")
	assert.Equal(t, true, test1)
	test2 := StringEndsWith(" ", "")
	assert.Equal(t, true, test2)
	test3 := StringEndsWith("abc", "c")
	assert.Equal(t, true, test3)
	test4 := StringEndsWith("banana", "ana")
	assert.Equal(t, true, test4)
	test5 := StringEndsWith("a", "z")
	assert.Equal(t, false, test5)
	test6 := StringEndsWith("", "t")
	assert.Equal(t, false, test6)
	test7 := StringEndsWith("$a = $b + 1", "+1")
	assert.Equal(t, false, test7)
	test8 := StringEndsWith("1oo", "100")
	assert.Equal(t, false, test8)
	test9 := StringEndsWith("    ", "   ")
	assert.Equal(t, true, test9)
	test10 := StringEndsWith("ails", "fails")
	assert.Equal(t, true, test10)
	test11 := StringEndsWith("abcabc", "bc")
	assert.Equal(t, true, test11)
	test12 := StringEndsWith("abc", "abc")
	assert.Equal(t, true, test12)
	test13 := StringEndsWith("samurai", "ai")
	assert.Equal(t, true, test13)
}

func TestCountByX(t *testing.T) {
	test1 := CountByX(2, 5)
	assert.Equal(t, []int{2, 4, 6, 8, 10}, test1)
	test2 := CountByX(1, 5)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, test2)
	test3 := CountByX(3, 5)
	assert.Equal(t, []int{3, 6, 9, 12, 15}, test3)
	test4 := CountByX(50, 5)
	assert.Equal(t, []int{50, 100, 150, 200, 250}, test4)
	test5 := CountByX(100, 5)
	assert.Equal(t, []int{100, 200, 300, 400, 500}, test5)
}
