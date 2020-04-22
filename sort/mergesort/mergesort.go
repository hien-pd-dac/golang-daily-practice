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
	for i := 0; i < resultLen; i++ {
		leftPopValue := popFirstElement(leftArr)
		rightPopValue := popFirstElement(rightArr)
		smallerMark := getSmaller(leftPopValue, rightPopValue)
		if smallerMark == -1 {
			log.Printf("error when 2 array are both empty in turn of loop: %d", i)
			return []int{}
		}
		if smallerMark == 0 {
			resultArr[i] = *leftPopValue
			leftArr = leftArr[1:len(leftArr)]
		} else {
			resultArr[i] = *rightPopValue
			rightArr = rightArr[1:len(rightArr)]
		}
	}
	log.Printf("end merge: leftArr=%v, rightArr=%v., result=%v", leftArr, rightArr, resultArr)
	return resultArr
}

// getSmaller return -1 if both of first and second are nil
// return 0 if first is larger
// rerturn 1 if second is larger
func getSmaller(first, second *int) (mark int) {
	if first == nil && second == nil {
		return -1
	}
	if first == nil {
		return 1
	}
	if second == nil {
		return 0
	}
	if *first < *second {
		return 0
	}
	return 1
}

func popFirstElement(arr []int) *int {
	if len(arr) == 0 {
		return nil
	}
	return &arr[0]
}
