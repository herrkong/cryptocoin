package hash

import (
	"testing"
)

func TestHash(t *testing.T) {
	type args struct {
		Input string
	}
	tests := []struct {
		name       string
		args       args
		wantOutPut string
		wantErr    bool
	}{
		// TODO: Add test cases.
		{
			name: "test sha256 hash",
			args: args{Input: "darwinkong"},
			wantOutPut: "25be9c79f92c5f8c56bfdd99f516712d424a0e8e28c976e4387723a787f2dd2d",
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOutPut, err := Hash(tt.args.Input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Hash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOutPut != tt.wantOutPut {
				t.Errorf("Hash() = %v, want %v", gotOutPut, tt.wantOutPut)
			}
		})
	}
}
