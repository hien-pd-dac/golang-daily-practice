package mergesort

import "log"

type Sorter interface {
	Sort(arr []int) []int
}

type sorter struct{}

func NewSorter() Sorter {
	return &sorter{}
}

func (*sorter) Sort(arr []int) []int {
	return mergeSort(arr)
}

func mergeSort(arr []int) []int {
	arrLen := len(arr)
	copyArr := make([]int, arrLen)
	copy(copyArr, arr)
	if arrLen <= 1 {
		return copyArr
	}
	midIdx := arrLen / 2
	leftArr := mergeSort(copyArr[:midIdx])
	rightArr := mergeSort(copyArr[midIdx:])
	return merge(leftArr, rightArr)
}

func merge(leftArr, rightArr []int) []int {
	log.Printf("start merge: leftArr=%v, rightArr=%v.", leftArr, rightArr)
	leftArrLen := len(leftArr)
	rightArrLen := len(rightArr)
	resultLen := leftArrLen + rightArrLen
	resultArr := make([]int, resultLen)
	resultIdx := 0
	for !(len(leftArr) == 0 && len(rightArr) == 0) {
		var smallerVal int
		if len(leftArr) == 0 {
			smallerVal, rightArr = popFirst(rightArr)
		} else if len(rightArr) == 0 || rightArr[0] > leftArr[0] {
			smallerVal, leftArr = popFirst(leftArr)
		} else {
			smallerVal, rightArr = popFirst(rightArr)
		}

		resultArr[resultIdx] = smallerVal
		resultIdx += 1
	}
	log.Printf("end merge: leftArr=%v, rightArr=%v., result=%v", leftArr, rightArr, resultArr)
	return resultArr
}

// popFirst pops the first element of an array, return this value and new array
func popFirst(source []int) (value int, result []int) {
	value, result = source[0], source[1:]
	return
}
