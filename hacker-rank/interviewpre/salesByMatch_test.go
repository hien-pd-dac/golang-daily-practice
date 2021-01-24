package interviewpre

import "testing"

func TestSockMerchant(t *testing.T) {
	type args struct {
		n  int32
		ar []int32
	}
	testCases := []struct {
		desc string
		args args
		want int32
	}{
		{
			desc: "Should return 0 pairs if n = 1.",
			args: args{
				n:  int32(1),
				ar: []int32{1},
			},
			want: int32(0),
		},
		{
			desc: "Should return 1 pair if have only 1 pair.",
			args: args{
				n:  int32(2),
				ar: []int32{1, 1},
			},
			want: int32(1),
		},
		{
			desc: "Should return 1 pair if have 1 pair and 1 sock.",
			args: args{
				n:  int32(3),
				ar: []int32{1, 2, 1},
			},
			want: int32(1),
		},
		{
			desc: "Should return 4 pairs if have 4 pairs",
			args: args{
				n:  int32(11),
				ar: []int32{1, 2, 1, 3, 2, 4, 4, 3, 5, 6, 7},
			},
			want: int32(4),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			got := sockMerchant(tc.args.n, tc.args.ar)
			if got != tc.want {
				t.Errorf("\n\tgot:  %d.\n\twant: %d.\n", got, tc.want)
			}

		})
	}
}
