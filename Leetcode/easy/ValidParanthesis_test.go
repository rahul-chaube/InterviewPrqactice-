package easy

import "testing"

func TestIsValid(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Positive",
			args: args{
				str: "[]()",
			},
			want: true,
		},
		{
			name: "Positive 1",
			args: args{
				str: "[()]()",
			},
			want: true,
		},
		{
			name: "Negative",
			args: args{
				str: "[(])",
			},
			want: false,
		},
		{
			name: "Negative",
			args: args{
				str: "[]{})",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValid(tt.args.str); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
