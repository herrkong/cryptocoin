package bubbleSort

import (
	"reflect"
	"testing"
)

func TestBubleSort(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "heap Sort",
			args: args{array: []int{6, 4, 7, 3, 1, 2}},
			want: []int{1, 2, 3, 4, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BubleSort(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BubleSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
