package selectSort

import (
	"reflect"
	"testing"
)

func TestSelectSort(t *testing.T) {
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
			name: "select sort",
			args: args{array: []int{5, 4, 3, 6, 2, 1}},
			want: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SelectSort(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
