package warmUp

import "testing"

func TestRepeatedString(t *testing.T) {
	type args struct {
		s string
		n int64
	}
	testCases := []struct {
		desc string
		args args
		want int64
	}{
		{
			desc: "",
			args: args{
				s: "aba",
				n: int64(10),
			},
			want: int64(7),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			got := repeatedString(tc.args.s, tc.args.n)
			if got != tc.want {
				t.Errorf("\n\tgot: %d.\n\twant: %d.\n", got, tc.want)
			}

		})
	}
}
