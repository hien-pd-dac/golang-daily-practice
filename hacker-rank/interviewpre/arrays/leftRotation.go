package arrays

// Complete the rotLeft function below.
func rotLeft(a []int32, d int32) []int32 {
	aLen := len(a)
	surplus := int(d) % aLen
	if surplus == 0 {
		return a
	}
	targetArr := make([]int32, 0, aLen)
	targetArr = append(targetArr, a[surplus:]...)
	targetArr = append(targetArr, a[:surplus]...)
	return targetArr
}
