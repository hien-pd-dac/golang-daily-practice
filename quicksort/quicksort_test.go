package quicksort

import "testing"

func TestQuickSort(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{
			name: "empty array",
			args: []int{},
			want: []int{},
		},
		{
			name: "only one element",
			args: []int{1},
			want: []int{1},
		},
		{
			name: "2 elements",
			args: []int{1, 2},
			want: []int{1, 2},
		},
		{
			name: "more than 2 elements & each element is not equal to the other",
			args: []int{2, 1, 3, 8, 0},
			want: []int{0, 1, 2, 3, 8},
		},
		{
			name: "more than 2 elements & each element is not equal to the other",
			args: []int{2, 1, 3, 8, 0, 3, 1, 7},
			want: []int{0, 1, 1, 2, 3, 3, 7, 8},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Quicksort(tt.args)
			if !isEqual(got, tt.want) {
				t.Errorf("\ngot: %v\nget: %v\n", got, tt.want)
			}
		})
	}
}

func isEqual(source []int, target []int) bool {
	sourceLen := len(source)
	if sourceLen != len(target) {
		return false
	}
	for i := 0; i < sourceLen; i++ {
		if source[i] != target[i] {
			return false
		}
	}
	return true
}
