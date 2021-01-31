package arrays

// arr is 6x6 2D-array
func hourglassSum(arr [][]int32) int32 {
	max := int32(-999999)
	for i := 1; i <= 4; i++ {
		for j := 1; j <= 4; j++ {
			sumOfCurHourGlass := sumOfHourGlass(arr, i, j)
			if sumOfCurHourGlass > max {
				max = sumOfCurHourGlass
			}
		}
	}
	return max
}

func sumOfHourGlass(arr [][]int32, xi, yj int) int32 {
	sum := int32(0)
	for i := xi - 1; i <= xi+1; i++ {
		if i == xi {
			sum = sum + arr[i][yj]
			continue
		}
		for j := yj - 1; j <= yj+1; j++ {
			sum = sum + arr[i][j]

		}
	}
	return sum
}
