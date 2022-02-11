package heapsort

import (
	"reflect"
	"testing"
)

func TestHeapSort(t *testing.T) {
	type args struct {
		Input []int
	}
	tests := []struct {
		name       string
		args       args
		wantOutPut []int
	}{
		// TODO: Add test cases.
		{	
			name: "heap Sort",
			args: args{Input: []int{6,4,7,3,1,2}},
			wantOutPut: []int{1,2,3,4,6,7},

		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutPut := HeapSort(tt.args.Input); !reflect.DeepEqual(gotOutPut, tt.wantOutPut) {
				t.Errorf("HeapSort() = %v, want %v", gotOutPut, tt.wantOutPut)
			}
		})
	}
}
