package quicksort

type Sorter interface {
	Sort(arr []int) []int
}

type sorter struct{}

func NewSorter() Sorter {
	return &sorter{}
}

func (s *sorter) Sort(arr []int) []int {
	arrLen := len(arr)
	copyArr := make([]int, arrLen)
	copy(copyArr, arr)
	sort(copyArr, 0, arrLen-1)
	return copyArr
}

func sort(arr []int, fromIdx int, toIdx int) {
	// fmt.Printf("start quicksort: arr=%v, fromIdx=%d, toIdx=%d.\n", arr[fromIdx:toIdx+1], fromIdx, toIdx)
	if fromIdx < 0 || toIdx < 0 || fromIdx > toIdx {
		// fmt.Printf("error: arr: %v, fromIdx:%d, toIdx:%d\n", arr[fromIdx:toIdx+1], fromIdx, toIdx)
		return
	}
	if (toIdx - fromIdx) < 1 {
		// fmt.Printf("end quicksort: arr=%v, fromIdx=%d, toIdx=%d.\n", arr[fromIdx:toIdx+1], fromIdx, toIdx)
		return
	}
	sortPartition := arr[fromIdx : toIdx+1]
	pivotIdx := getPivot(sortPartition)
	newPivotIdx := splitByPivot(sortPartition, pivotIdx)
	if newPivotIdx > 0 {
		sort(sortPartition, 0, newPivotIdx-1)
	}
	if lenPart := len(sortPartition); newPivotIdx < lenPart-1 {
		sort(sortPartition, newPivotIdx+1, lenPart-1)
	}
	// fmt.Printf("end quicksort: arr=%v, fromIdx=%d, toIdx=%d.\n", arr[fromIdx:toIdx+1], fromIdx, toIdx)
}

// getPivot returns the index of the selected pivot.
func getPivot(arr []int) int {
	return len(arr) / 2
}

// if not really have LeftPart or RightPart, the corresponse Indx will be -1
func splitByPivot(arr []int, pivotIdx int) (newPivotIdx int) {
	lastIdx := len(arr) - 1
	swap(arr, pivotIdx, lastIdx)
	leftBoundIdx := 0
	rightBoundIdx := lastIdx
	for {
		leftBoundIdx = getIdxOfTheEqualOrGreaterThanPivot(arr, leftBoundIdx)
		rightBoundIdx = getIdxOfTheLessThanPivot(arr, rightBoundIdx)
		if rightBoundIdx < leftBoundIdx {
			break
		}
		swap(arr, leftBoundIdx, rightBoundIdx)
	}
	if leftBoundIdx != lastIdx {
		swap(arr, leftBoundIdx, lastIdx)
	}

	return leftBoundIdx
}

func swap(arr []int, sourceIdx, destIdx int) {
	arr[sourceIdx], arr[destIdx] = arr[destIdx], arr[sourceIdx]
}

// returns the index of the element is greater than or equal pivot.
// returns the last index of this array (the pivot's index) if does not exist that element.
func getIdxOfTheEqualOrGreaterThanPivot(arr []int, fromIdx int) (resultIdx int) {
	arrLen := len(arr)
	pivot := arr[arrLen-1]
	resultIdx = arrLen - 1
	for i := fromIdx; i < arrLen; i++ {
		if arr[i] >= pivot {
			resultIdx = i
			break
		}
	}
	return
}

// returns the index of the element is less than pivot.
// returns -1 if does not exist that element.
func getIdxOfTheLessThanPivot(arr []int, fromReverseIdx int) int {
	arrLen := len(arr)
	pivot := arr[arrLen-1]
	for i := fromReverseIdx; i >= 0; i-- {
		if arr[i] < pivot {
			return i
		}
	}
	return -1
}
