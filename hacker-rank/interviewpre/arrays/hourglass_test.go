package arrays

import "testing"

func TestHourglass(t *testing.T) {
	testCases := []struct {
		desc string
		args [][]int32
		want int32
	}{
		{
			desc: "",
			args: [][]int32{
				{1, 1, 1, 0, 0, 0},
				{0, 1, 0, 0, 0, 0},
				{1, 1, 1, 0, 0, 0},
				{0, 0, 2, 4, 4, 0},
				{0, 0, 0, 2, 0, 0},
				{0, 0, 1, 2, 4, 0},
			},
			want: int32(19),
		},
		{
			desc: "",
			args: [][]int32{
				{9, 0, 0, 0, 0, -2},
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{1, 0, 0, 0, 0, 1},
			},
			want: int32(9),
		},
		{
			desc: "",
			args: [][]int32{
				{1, 1, 1, 0, 0, 0},
				{0, 1, 0, 0, 0, 0},
				{1, 1, 1, 0, 0, 0},
				{0, 9, 2, -4, -4, 0},
				{0, 0, 0, -2, 0, 0},
				{1, 0, -1, -2, -4, 0},
			},
			want: int32(13),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			got := hourglassSum(tc.args)
			if got != tc.want {
				t.Errorf("\n\tgot: %d.\n\twant: %d.\n", got, tc.want)
			}
		})
	}
}
