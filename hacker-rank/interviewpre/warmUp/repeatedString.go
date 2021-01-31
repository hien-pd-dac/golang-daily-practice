package warmUp

func repeatedString(s string, n int64) int64 {
	sLen := int64(len(s))
	divided := n / sLen
	du := n % sLen
	numOfAInString := getTheNumberOfALetterInString(s, sLen-1)
	numOfAInStringToDuIdx := getTheNumberOfALetterInString(s, du-1)
	return divided*int64(numOfAInString) + int64(numOfAInStringToDuIdx)
}

func getTheNumberOfALetterInString(s string, idx int64) int {
	count := 0
	for i := int64(0); i <= idx; i++ {
		if s[i] == 'a' {
			count = count + 1
		}
	}
	return count

}
