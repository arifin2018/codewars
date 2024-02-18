package Go

// https://www.codewars.com/kata/5513795bd3fafb56c200049e/train/go
func CountByX(x, n int) []int {
	var result = []int{x}
	for i := 1; i < n; i++ {
		result = append(result, result[i-1]+x)
	}
	return result
}
