package factor

type Detector interface {
	Factors() []int
}

type detector struct {
	number int
}

func NewDetector(number int) Detector {
	return &detector{
		number: number,
	}
}

func (d *detector) Factors() []int {
	factors := []int{}
	number := d.number
	if number < 2 {
		return []int{}
	}
	for i := 2; number > 1 && i <= number; i++ {
		if number%i == 0 {
			var dividedCount = 0
			number, dividedCount = divideFor(number, i)
			factors = appendWithCount(factors, i, dividedCount)
		}
	}
	return factors
}

func divideFor(number, dividen int) (newNum, count int) {
	count = 0
	for number%dividen == 0 {
		number = number / dividen
		count = count + 1
	}
	return number, count
}

func appendWithCount(dest []int, element, count int) []int {
	copyDest := make([]int, len(dest), cap(dest)+count)
	copy(copyDest, dest)
	for i := 0; i < count; i++ {
		copyDest = append(copyDest, element)
	}
	return copyDest
}
