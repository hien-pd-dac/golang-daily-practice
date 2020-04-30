package binary

type Searcher interface {
	Search(arr []int, element int) int
}

type searcher struct{}

func New() Searcher {
	return &searcher{}
}

func (b *searcher) Search(arr []int, element int) int {
	lenArr := len(arr)
	if lenArr == 0 {
		return -1
	}
	return search(arr, element, 0, lenArr-1)
}

func search(arr []int, element, fromIdx, toIdx int) int {
	if toIdx-fromIdx == 0 {
		if arr[fromIdx] == element {
			return fromIdx
		}
		return -1
	}
	middleIdx := (toIdx + fromIdx) / 2
	if leftSearchResult := search(arr, element, fromIdx, middleIdx); leftSearchResult != -1 {
		return leftSearchResult
	}
	return search(arr, element, middleIdx+1, toIdx)
}
