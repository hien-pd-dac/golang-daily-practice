package hashmap

//https://www.hackerrank.com/challenges/ctci-ransom-note/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=dictionaries-hashmaps
// func checkMagazine(magazine []string, note []string) {
// fmt.Printf("%s\n", mainCheckMagazine(magazine, note))
// }

func mainCheckMagazine(magazine []string, note []string) string {
	lenMagazine := len(magazine)
	lenNote := len(note)
	if lenMagazine < lenNote {
		return "No"
	}
	magazineDict := genDictFromStrArr(magazine)
	noteDict := genDictFromStrArr(note)

	for k, v := range noteDict {
		numInMagazine, ok := magazineDict[k]
		if !ok {
			return "No"
		}

		if v > numInMagazine {
			return "No"
		}
	}
	return "Yes"
}

func genDictFromStrArr(arr []string) map[string]int {
	dict := make(map[string]int)
	for _, value := range arr {
		_, ok := dict[value]
		if !ok {
			dict[value] = 1
			continue
		}
		dict[value] = dict[value] + 1
	}
	return dict
}
