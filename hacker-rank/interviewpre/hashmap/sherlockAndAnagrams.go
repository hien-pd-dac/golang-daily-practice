package hashmap

type dictSherlock map[string]int32

type algoSherlock struct {
	s    string
	dict dictSherlock
}

// https://www.hackerrank.com/challenges/sherlock-and-anagrams/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=dictionaries-hashmaps
func sherlockAndAnagrams(s string) int32 {
	al := &algoSherlock{
		s:    s,
		dict: dictSherlock(make(map[string]int32)),
	}

	for l := 1; l <= len(s); l++ {
		al.writeToDictWithSpecifiedLen(l)
	}

	return al.countPairs()
}

func (al *algoSherlock) writeToDictWithSpecifiedLen(l int) {
	for i := 0; i+l-1 < len(al.s); i++ {
		strWithLen := al.s[i : i+l]
		al.writeToDict(i, strWithLen)
	}
	return
}

func (al *algoSherlock) writeToDict(idx int, s string) {
	s = sortByAlphabet(s)
	if _, ok := al.dict[s]; !ok {
		al.dict[s] = 1
		return
	}
	al.dict[s] = al.dict[s] + 1
	return
}

func sortByAlphabet(s string) string {
	return sort(s)
}

func (al *algoSherlock) countPairs() int32 {
	count := int32(0)
	for _, v := range al.dict {
		if v < 2 {
			continue
		}
		count += c2Of(int(v))
	}
	return count
}

func c2Of(n int) int32 {
	return int32(n * (n - 1) / 2)
}

func sort(s string) string {
	sLen := len(s)
	for i := 0; i < sLen-1; i++ {
		minEleIdx := min(i, s)
		if minEleIdx != i {
			s = swap(s, i, minEleIdx)
		}
	}
	return s
}

func min(idx int, s string) int {
	minIdx := idx
	for i := idx; i < len(s); i++ {
		if s[i] < s[minIdx] {
			minIdx = i
		}
	}
	return minIdx
}

func swap(s string, i, j int) string {
	slice := []byte(s)
	tmp := slice[i]
	slice[i] = slice[j]
	slice[j] = tmp
	return string(slice)
}
