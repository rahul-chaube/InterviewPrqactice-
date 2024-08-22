package easy

import "testing"

func TestFindCommonPrefix(t *testing.T) {
	type args struct {
		str []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Positive",
			args: args{
				str: []string{"flo", "flower", "flam"},
			},
			want: "fl",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findCommonPrefix(tt.args.str); got != tt.want {
				t.Errorf("findCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
