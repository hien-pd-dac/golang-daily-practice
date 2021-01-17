package rank

func mostBalancedPartition(parent []int32, files_size []int32) int32 {
	n := int32(len(parent))
	parentToChildMap := convertToParentChildMap(parent)

	sumSizeOfEachNodeMap := make(map[int32]int32)

	calculateContentSizeEachNode(sumSizeOfEachNodeMap, parentToChildMap, files_size, 0)

	sumSizeOfAll := sumAll(files_size)
	minDiff := int32(999999)
	for i := int32(1); i < n; i++ {
		diffPartition := sumSizeOfAll - 2*sumSizeOfEachNodeMap[i]
		if diffPartition < minDiff {
			minDiff = diffPartition
		}

	}
	return minDiff

}

func convertToParentChildMap(parent []int32) map[int32][]int32 {
	parentToChildMap := make(map[int32][]int32)
	for i := int32(0); i < int32(len(parent)); i++ {
		parentNode := parent[i]
		parentToChildMap[parentNode] = append(parentToChildMap[parentNode], i)
	}
	return parentToChildMap

}

func calculateContentSizeEachNode(
	sumSizeEachNodeMap map[int32]int32,
	parentToChildMap map[int32][]int32,
	files_size []int32,
	curNodeIdx int32) int32 {
	if len(parentToChildMap[curNodeIdx]) == 0 {
		sumSizeEachNodeMap[curNodeIdx] = files_size[curNodeIdx]
		return sumSizeEachNodeMap[curNodeIdx]
	}
	for _, child := range parentToChildMap[curNodeIdx] {
		sumSizeEachNodeMap[curNodeIdx] = sumSizeEachNodeMap[curNodeIdx] + calculateContentSizeEachNode(sumSizeEachNodeMap, parentToChildMap, files_size, child)
	}
	return sumSizeEachNodeMap[curNodeIdx]
}

func sumAll(files_size []int32) int32 {
	sum := int32(0)
	for _, size := range files_size {
		sum = sum + size
	}
	return sum
}
