package chapter1

import "testing"

func TestConvert(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Positive 1",
			args: args{
				num: 10,
			},
		},
		{
			name: "Positive 2",
			args: args{
				num: 9,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Convert(tt.args.num)
		})
	}
}
