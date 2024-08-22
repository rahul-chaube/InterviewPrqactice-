package easy

import "testing"

func TestDiffString(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Positive 1",
			args: args{
				str1: "abcde",
				str2: "abcd",
			},
			want: "e",
		},
		{
			name: "Positive 2",
			args: args{
				str1: "abcde",
				str2: "abcf",
			},
			want: "def",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DiffString(tt.args.str1, tt.args.str2); got != tt.want {
				t.Errorf("DiffString() = %v, want %v", got, tt.want)
			}
		})
	}
}
