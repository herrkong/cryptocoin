package bb

import "testing"

func TestBb(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bb(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Bb() = %v, want %v", got, tt.want)
			}
		})
	}
}
