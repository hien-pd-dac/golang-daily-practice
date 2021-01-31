package warmUp

func jumpingOnClouds(c []int32) int32 {
	// init curIdx, curSteps.
	cLength := len(c)
	curIdx := 0
	jumpedSteps := int32(0)
	// loop until the end.
	for {
		nextJump := getTheNextJump(curIdx, c)
		jumpedSteps = jumpedSteps + int32(1)
		curIdx = curIdx + int(nextJump)
		if curIdx == cLength-1 {
			break
		}
	}
	return jumpedSteps
}

func getTheNextJump(curIdx int, c []int32) int32 {
	if curIdx+2 < len(c) && c[curIdx+2] == int32(0) {
		return int32(2)
	}
	return int32(1)
}
