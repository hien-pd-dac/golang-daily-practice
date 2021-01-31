package warmUp

import "testing"

func TestCountingValleys(t *testing.T) {
	type args struct {
		steps int32
		path  string
	}
	testCases := []struct {
		desc string
		args args
		want int32
	}{
		{
			desc: "Should return 0 when having only 1 mountain.",
			args: args{
				steps: 2,
				path:  "UD",
			},
			want: int32(0),
		},
		{
			desc: "Should return 0 when having only mountains.",
			args: args{
				steps: 4,
				path:  "UDUD",
			},
			want: int32(0),
		},
		{
			desc: "Should return 1 when having 1 valley.",
			args: args{
				steps: 2,
				path:  "DU",
			},
			want: int32(1),
		},
		{
			desc: "Should return 2 when having 2 valleys.",
			args: args{
				steps: 6,
				path:  "DUUDDU",
			},
			want: int32(2),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			got := countingValleys(tc.args.steps, tc.args.path)
			if got != tc.want {
				t.Errorf("\n\tgot:  %d.\n\twant: %d.\n", got, tc.want)
			}

		})
	}
}
