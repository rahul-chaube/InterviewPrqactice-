package easy

import "testing"

func TestIsPolidromNumber(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Positive 1",
			args: args{number: 121},
			want: true,
		},
		{
			name: "Positive 2",
			args: args{number: 123},
			want: false,
		},
		{
			name: "Positive 2",
			args: args{number: -121},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPolidromNumber(tt.args.number); got != tt.want {
				t.Errorf("IsPolidromNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
