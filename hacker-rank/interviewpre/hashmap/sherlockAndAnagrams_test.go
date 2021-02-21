package hashmap

import "testing"

func TestSherlockAndAnagrams(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "",
			args: args{
				s: "ifailuhkqq",
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				s: "kkkk",
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sherlockAndAnagrams(tt.args.s); got != tt.want {
				t.Errorf("sherlockAndAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
