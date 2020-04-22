package mergesort

import (
	"testing"

	"github.com/hienpdbk/daily-practice/slice"
)

type testData struct {
	args []int
	want []int
}

func TestSort(t *testing.T) {
	tests := []struct {
		name      string
		testdatas []testData
	}{
		{
			name: "empty array",
			testdatas: []testData{
				testData{
					args: []int{},
					want: []int{},
				},
			},
		},
		{
			name: "one element array",
			testdatas: []testData{
				testData{
					args: []int{1},
					want: []int{1},
				},
			},
		},
		{
			name: "more than or equal two elements array",
			testdatas: []testData{
				testData{
					args: []int{1, 2},
					want: []int{1, 2},
				},
				testData{
					args: []int{1, 2, 3},
					want: []int{1, 2, 3},
				},
				testData{
					args: []int{3, 2, 1},
					want: []int{1, 2, 3},
				},
				testData{
					args: []int{1, 2, 3, 4, 5, 4, 3, 2, 1},
					want: []int{1, 1, 2, 2, 3, 3, 4, 4, 5},
				},
			},
		},
	}

	mergeSort := NewSorter()
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			for _, data := range tc.testdatas {
				got := mergeSort.Sort(data.args)
				if !slice.IsSame(got, data.want) {
					t.Errorf("\ngot: %v\nwant: %v\n", got, data.want)
				}
			}

		})
	}
}

func TestMergeWithBothEmptyArray(t *testing.T) {
	if len(merge([]int{}, []int{})) != 0 {
		t.Errorf("Failed with merge 2 empty array")
	}
}
