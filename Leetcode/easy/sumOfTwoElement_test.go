package easy

import (
	"reflect"
	"testing"
)

func TestSumOfTwoNumber(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test case 1",
			args: args{
				arr:    []int{3, 3, 8, 10, 3, 5, 7},
				target: 11,
			},
			want: []int{1, 2},
		},
		{
			name: "Test case 2",
			args: args{
				arr:    []int{3, 3, 8, 10, 3, 5, 7},
				target: 18,
			},
			want: []int{2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumOfTwoNumber(tt.args.arr, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumOfTwoNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
