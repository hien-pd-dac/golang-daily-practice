package interviewpre

const (
	UpStep   = "U"
	DownStep = "D"
)

func countingValleys(steps int32, path string) int32 {
	count := int32(0)
	curHeight := 0
	for i := 0; i < int(steps); i++ {
		if string(path[i]) == UpStep {
			curHeight = curHeight + 1
			if curHeight == 0 {
				count = count + 1
			}
			continue
		}
		curHeight = curHeight - 1
	}
	return count
}
