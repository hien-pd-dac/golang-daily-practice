package arrays

import (
	"fmt"
	"strconv"
)

// https://www.hackerrank.com/challenges/new-year-chaos/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=arrays
func minimumBribes(q []int32) {
	fmt.Printf("%s\n", mainMinimumBribes(q))
}

func mainMinimumBribes(q []int32) string {
	bribes := 0
	toIdx := len(q) - 1
	for toIdx > 0 {
		newQ, plusBribes := arrangeToTargetPosition(q, toIdx)
		if plusBribes == -1 {
			return "Too chaotic"
		}
		bribes = bribes + plusBribes
		q = newQ
		toIdx = toIdx - 1
	}
	return strconv.Itoa(bribes)
}

// if too chaotic, return (nil, -1)
func arrangeToTargetPosition(q []int32, targetIdx int) (newQ []int32, bribes int) {
	idx := getCurIdxOfTargetValue(int32(targetIdx+1), q)
	if idx == -1 {
		return nil, -1
	}
	bribes = targetIdx - idx
	if bribes == 0 {
		return q, 0
	}

	moveValue := q[idx]
	for i := idx; i < targetIdx; i++ {
		q[i] = q[i+1]
	}
	q[targetIdx] = moveValue
	return q, bribes
}

func getCurIdxOfTargetValue(val int32, q []int32) int {
	targetIdx := int(val - 1)
	for i := targetIdx; i >= targetIdx-2 && i >= 0; i-- {
		if val == q[i] {
			return i
		}
	}
	return -1
}
