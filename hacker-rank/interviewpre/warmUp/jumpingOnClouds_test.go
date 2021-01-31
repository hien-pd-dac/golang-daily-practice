package warmUp

import "testing"

func TestJumpingOnClouds(t *testing.T) {
	testCases := []struct {
		desc   string
		clouds []int32
		want   int32
	}{
		{
			desc:   "Should return 1 step for the shortest path with 2 clouds.",
			clouds: []int32{0, 0},
			want:   int32(1),
		},
		{
			desc:   "Should return the correct minimum number of jumps required.",
			clouds: []int32{0, 0, 1, 0, 0, 1, 0},
			want:   int32(4),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			got := jumpingOnClouds(tc.clouds)
			if got != tc.want {
				t.Errorf("\n\tgot:  %d.\n\twant: %d.\n", got, tc.want)
			}

		})
	}
}
