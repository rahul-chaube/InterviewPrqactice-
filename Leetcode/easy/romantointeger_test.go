package easy

import "testing"

func TestRomanToInteger(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "First Case 1",
			args: args{
				s: "III",
			},
			want: 3,
		},
		{
			name: "First Case 2",
			args: args{
				s: "IV",
			},
			want: 4,
		},
		{
			name: "First Case 3",
			args: args{
				s: "LVIII",
			},
			want: 58,
		},
		{
			name: "First Case 3",
			args: args{
				s: "MCMXCIV",
			},
			want: 1994,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RomanToInteger(tt.args.s); got != tt.want {
				t.Errorf("RomanToInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}
