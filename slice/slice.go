package slice

func IsSame(source []int, target []int) bool {
	sourceLen := len(source)
	if sourceLen != len(target) {
		return false
	}
	for i := 0; i < sourceLen; i++ {
		if source[i] != target[i] {
			return false
		}
	}
	return true
}
