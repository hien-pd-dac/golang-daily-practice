package arrays

import "testing"

func TestMinimumSwap(t *testing.T) {
	testCases := []struct {
		desc string
		args []int
		want int
	}{
		// {
		// 	desc: "Should return 0 swaps with 1 element-array.",
		// 	args: []int{1},
		// 	want: 0,
		// },
		// {
		// 	desc: "Should return 0 swaps with >1 sorted element-array.",
		// 	args: []int{1, 2, 3},
		// 	want: 0,
		// },
		{
			desc: "Should return 1 swap with element-array.",
			args: []int{2, 1, 3},
			want: 1,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			if got := minimumSwaps(tc.args); got != tc.want {
				t.Errorf("\n\tgot:  %v\n\twant: %v\n", got, tc.want)
			}
		})
	}
}
