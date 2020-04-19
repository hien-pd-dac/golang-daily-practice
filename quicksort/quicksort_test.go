package quicksort

import (
	"testing"

	"github.com/hienpdbk/daily-practice/slice"
)

type testData struct {
	args []int
	want []int
}

func TestQuickSort(t *testing.T) {
	tests := []struct {
		name  string
		datas []testData
	}{
		{
			name: "empty array",
			datas: []testData{
				testData{
					args: []int{},
					want: []int{},
				},
			},
		},
		{
			name: "only one element",
			datas: []testData{
				testData{
					args: []int{1},
					want: []int{1},
				},
			},
		},
		{
			name: "2 elements",
			datas: []testData{
				testData{
					args: []int{1, 2},
					want: []int{1, 2},
				},
				testData{
					args: []int{2, 1},
					want: []int{1, 2},
				},
			},
		},
		{
			name: "more than 2 elements",
			datas: []testData{
				testData{
					args: []int{1, 2, 3, 4},
					want: []int{1, 2, 3, 4},
				},
				testData{
					args: []int{5, 4, 3, 2},
					want: []int{2, 3, 4, 5},
				},
				testData{
					args: []int{3, 8, 9, 1, 7},
					want: []int{1, 3, 7, 8, 9},
				},
				testData{
					args: []int{3, 8, 3, 1, 7},
					want: []int{1, 3, 3, 7, 8},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			for _, data := range tc.datas {
				sorter := NewSorter()
				got := sorter.Sort(data.args)
				if !slice.IsSame(got, data.want) {
					t.Errorf("\ngot: %v\nget: %v\n", got, data.want)
				}
			}
		})
	}
}
