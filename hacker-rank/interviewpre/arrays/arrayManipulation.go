package arrays

func arrayManipulation(n int32, queries [][]int32) int64 {
	origin := make([]int64, n)
	max := int64(0)

	for _, query := range queries {
		fromIdx := query[0] - 1
		toIdx := query[1] - 1
		plus := query[2]
		manipulate(origin, fromIdx, toIdx, plus)
	}

	cur := int64(0)
	for _, v := range origin {
		cur += v
		if cur > max {
			max = cur
		}
	}

	return max
}

func manipulate(arr []int64, fromIdx, toIdx, plus int32) {
	arr[fromIdx] += int64(plus)
	if toIdx+1 < int32(len(arr)) {
		arr[toIdx+1] -= int64(plus)
	}
	return
}
