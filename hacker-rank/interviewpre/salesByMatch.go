package interviewpre

func sockMerchant(n int32, ar []int32) int32 {
	sockIDAndTheNumOfSocksMap := genSockIDAndTheNumOfSocksMap(n, ar)
	return countSockPairs(sockIDAndTheNumOfSocksMap)
}

func genSockIDAndTheNumOfSocksMap(n int32, ar []int32) map[int32]int32 {
	sockMap := make(map[int32]int32)
	for _, val := range ar {
		sockMap[val] = sockMap[val] + 1
	}
	return sockMap
}

func countSockPairs(sockMap map[int32]int32) int32 {
	numOfPairs := int32(0)
	for _, val := range sockMap {
		numOfPairs = numOfPairs + val/2
	}
	return numOfPairs
}
