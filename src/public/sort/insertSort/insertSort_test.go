package insertSort

import (
	"reflect"
	"testing"
)

func TestInsertSort(t *testing.T) {
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
			name:       "insert sort",
			args:       args{Input: []int{5, 4, 3, 6, 2, 1}},
			wantOutPut: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutPut := InsertSort(tt.args.Input); !reflect.DeepEqual(gotOutPut, tt.wantOutPut) {
				t.Errorf("InsertSort() = %v, want %v", gotOutPut, tt.wantOutPut)
			}
		})
	}
}
