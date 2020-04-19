package factor

import (
	"testing"

	"github.com/hienpdbk/daily-practice/slice"
)

func TestFactor(t *testing.T) {
	tests := []struct {
		name string
		args int
		want []int
	}{
		{
			name: "Expect empty with args: 0.",
			args: 0,
			want: []int{},
		},
		{
			name: "Expect empty with args: 1.",
			args: 1,
			want: []int{},
		},
		{
			name: "Expect not empty with prime args: 7.",
			args: 7,
			want: []int{7},
		},
		{
			name: "Expect not empty with non-prime args: 84.",
			args: 84,
			want: []int{2, 2, 3, 7},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			detector := NewDetector(tc.args)
			got := detector.Factors()
			if !slice.IsSame(got, tc.want) {
				t.Errorf("\ngot: %v.\nwant: %v.\n", got, tc.want)
			}
		})
	}

}
