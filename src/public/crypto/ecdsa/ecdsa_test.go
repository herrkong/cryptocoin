package ecdsa

import "testing"

func TestExample(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "test ecdsa",
			args: args{msg: "darwin"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Example(tt.args.msg)
		})
	}
}
