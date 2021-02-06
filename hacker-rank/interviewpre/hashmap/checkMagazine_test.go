package hashmap

import "testing"

func TestCheckMagazine(t *testing.T) {
	type args struct {
		magazine []string
		note     []string
	}
	testCases := []struct {
		desc string
		args args
		want string
	}{
		{
			desc: "Should return 'No' When the number of element note[i] in note is larger than in magazine. ",
			args: args{
				magazine: []string{"a", "b"},
				note:     []string{"a", "a"},
			},
			want: "No",
		},
		{
			desc: "Should return 'No' When existing an element in note that not exsit in magazine. ",
			args: args{
				magazine: []string{"a", "b"},
				note:     []string{"a", "c"},
			},
			want: "No",
		},
		{
			desc: "Should return 'No' When len(note) > len(magazine).",
			args: args{
				magazine: []string{"a", "b"},
				note:     []string{"a", "b", "c"},
			},
			want: "No",
		},
		{
			desc: "Should return 'Yes' When OK.",
			args: args{
				magazine: []string{"a", "b", "c"},
				note:     []string{"a", "c"},
			},
			want: "Yes",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if got := mainCheckMagazine(tC.args.magazine, tC.args.note); got != tC.want {
				t.Errorf("\n\tgot: %v\n\twant: %v\n", got, tC.want)
			}
		})
	}
}
