package binary

import "testing"

type datatest struct {
	args    []int
	element int
	want    int
}

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name string
		data []datatest
	}{
		{
			name: "empty array",
			data: []datatest{
				datatest{
					args:    []int{},
					element: 5,
					want:    -1,
				},
			},
		},
		{
			name: "one element",
			data: []datatest{
				datatest{
					args:    []int{1},
					element: 5,
					want:    -1,
				},
				datatest{
					args:    []int{5},
					element: 5,
					want:    0,
				},
			},
		},
		{
			name: "more than 2-elements",
			data: []datatest{
				datatest{
					args:    []int{1, 2},
					element: 5,
					want:    -1,
				},
				datatest{
					args:    []int{1, 5},
					element: 5,
					want:    1,
				},
				datatest{
					args:    []int{1, 2, 3, 4, 5, 6, 7},
					element: 4,
					want:    3,
				},
				datatest{
					args:    []int{1, 2, 3, 4, 5, 6, 7},
					element: 7,
					want:    6,
				},
				datatest{
					args:    []int{1, 2, 3, 4, 5, 6, 7},
					element: 1,
					want:    0,
				},
			},
		},
	}

	searcher := New()

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			for _, data := range tc.data {
				got := searcher.Search(data.args, data.element)
				if got != data.want {
					t.Errorf("\ngot: %d\nwant: %d\n", got, data.want)
				}
			}
		})
	}
}
