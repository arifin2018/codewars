package Go

func CheckForFactorMySelf(base int, factor int) bool {
	if base%factor == 0 {
		return true
	} else {
		return false
	}
}
func CheckForFactor(base int, factor int) bool {
	return base%factor == 0
}
