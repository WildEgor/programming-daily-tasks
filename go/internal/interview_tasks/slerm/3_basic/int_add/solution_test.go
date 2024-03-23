package int_add

import "testing"

func Test_add(t *testing.T) {
	tests := []struct {
		name string
		a    int32
		want int32
	}{
		{"1", 0, 1},
		{"2", 2147483647, -2147483648},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add(tt.a); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}
