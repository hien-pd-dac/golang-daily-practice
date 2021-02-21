package hashmap

import "testing"

func TestTwoStrings(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				s1: "and",
				s2: "art",
			},
			want: "YES",
		},
		{
			name: "",
			args: args{
				s1: "be",
				s2: "cat",
			},
			want: "NO",
		},
		{
			name: "",
			args: args{
				s1: "bbee",
				s2: "cate",
			},
			want: "YES",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoStrings(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("twoStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
