package factorial

import "testing"

func Test_factorial(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"1", 0, 1},
		{"2", 1, 1},
		{"3", 2, 2},
		{"4", 3, 6},
		{"5", 4, 24},
		{"6", 5, 120},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := factorial(tt.n); got != tt.want {
				t.Errorf("factorial() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_factorial2(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"1", 0, 1},
		{"2", 1, 1},
		{"3", 2, 2},
		{"4", 3, 6},
		{"5", 4, 24},
		{"6", 5, 120},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := factorial2(tt.n); got != tt.want {
				t.Errorf("factorial2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_S_String(t *testing.T) {
	s := &S{1, 2}
	if got := s.String(); got != "1" {
		t.Errorf("S.String() = %v, want %v", got, "1")
	}
}
