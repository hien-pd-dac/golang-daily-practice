package arrays

import "testing"

func TestRotLeft(t *testing.T) {
	type args struct {
		a []int32
		d int32
	}
	testCases := []struct {
		desc string
		args args
		want []int32
	}{
		{
			desc: "Normal test.",
			args: args{
				a: []int32{1, 2, 3, 4, 5},
				d: int32(4),
			},
			want: []int32{5, 1, 2, 3, 4},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			got := rotLeft(tc.args.a, tc.args.d)
			if !equal(got, tc.want) {
				t.Errorf("\n\tgot: %v.\n\twant: %v.\n", got, tc.want)
			}

		})
	}
}

func equal(origin, dest []int32) bool {
	originLen := len(origin)
	destLen := len(dest)
	if originLen != destLen {
		return false
	}
	for i := 0; i < originLen; i++ {
		if origin[i] != dest[i] {
			return false
		}
	}
	return true
}
