package hashmap

func twoStrings(s1 string, s2 string) string {
	dict1 := make(map[uint8]bool)
	dict2 := make(map[uint8]bool)
	lenS1 := len(s1)
	lenS2 := len(s2)
	for i := 0; i < lenS1; i++ {
		if dict1[s1[i]] {
			continue
		}
		dict1[s1[i]] = true
		if dict2[s1[i]] {
			return "YES"
		}
		for j := 0; j < lenS2; j++ {
			dict2[s2[j]] = true
			if dict1[s2[j]] {
				return "YES"
			}
		}
	}
	return "NO"
}
