package arrays

func minimumSwaps(arr []int) int {
	valToKeyMap := genValToKeyFromArr(arr)
	alg := algorithm{
		arr:         arr,
		valToKeyMap: valToKeyMap,
	}
	_, swaps := alg.minimumSwaps(len(arr) - 1)
	return swaps
}

type algorithm struct {
	arr         []int
	valToKeyMap map[int]int
}

func genValToKeyFromArr(arr []int) map[int]int {
	valToKeyMap := make(map[int]int)
	for i, v := range arr {
		valToKeyMap[v] = i
	}
	return valToKeyMap
}

func (a algorithm) minimumSwaps(toIdx int) (al algorithm, swaps int) {
	if toIdx == 0 {
		return a, 0
	}
	maxEle := toIdx + 1
	maxEleIdx := a.valToKeyMap[maxEle]
	if maxEleIdx == toIdx {
		return a.minimumSwaps(toIdx - 1)
	}
	a = a.swap(maxEleIdx, toIdx)
	al, swaps = a.minimumSwaps(toIdx - 1)
	return al, swaps + 1
}

func (a algorithm) swap(originIdx, destIdx int) algorithm {
	originVal := a.arr[originIdx]
	destVal := a.arr[destIdx]
	a.arr[originIdx] = destVal
	a.arr[destIdx] = originVal
	a.valToKeyMap[destVal] = originIdx
	a.valToKeyMap[originVal] = destIdx
	return a
}
