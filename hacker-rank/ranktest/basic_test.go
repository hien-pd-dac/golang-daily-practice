package ranktest

import "testing"

func TestMostBalancedPartition(t *testing.T) {
	type args struct {
		parent     []int32
		files_size []int32
	}
	testCases := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "Should return the most balanced partition.",
			args: args{
				parent:     []int32{-1, 0, 0, 1, 1},
				files_size: []int32{3, 2, 6, 1, 3},
			},
			want: int32(3),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			got := mostBalancedPartition(tC.args.parent, tC.args.files_size)
			if got != tC.want {
				t.Fatalf("\n\tgot : %v.\n\twant: %v.\n", got, tC.want)
			}

		})
	}
}
