package arrays

import (
	"testing"
)

func TestMinimumBribes(t *testing.T) {
	testCases := []struct {
		desc string
		args []int32
		want string
	}{
		{
			desc: "normalTest1",
			args: []int32{2, 1, 5, 3, 4},
			want: "3",
		},
		{
			desc: "normalTest2",
			args: []int32{1, 2, 5, 3, 7, 8, 6, 4},
			want: "7",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			got := mainMinimumBribes(tc.args)
			if got != tc.want {
				t.Errorf("\n\tgot: %s.\n\twant: %s.\n", got, tc.want)
			}

		})
	}
}
