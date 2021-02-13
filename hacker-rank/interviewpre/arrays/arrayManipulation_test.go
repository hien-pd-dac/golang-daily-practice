package arrays

import (
	"testing"
)

func TestArrayManipulation(t *testing.T) {
	type args struct {
		n       int32
		queries [][]int32
	}
	testCases := []struct {
		desc string
		args args
		want int64
	}{
		{
			desc: "",
			args: args{
				n: int32(10),
				queries: [][]int32{
					{1, 5, 3},
					{4, 8, 7},
					{6, 9, 1},
				},
			},
			want: int64(10),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if got := arrayManipulation(tC.args.n, tC.args.queries); got != tC.want {
				t.Errorf("\n\tgot:  %v\n\twant: %v\n", got, tC.want)
			}

		})
	}
}
