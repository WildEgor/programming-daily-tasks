package __3_4

import "testing"

func Test_distance(t *testing.T) {
	tests := []struct {
		name  string
		lat1  float32
		lon1  float32
		lat2  float32
		lon2  float32
		want  float32
		want1 float32
	}{
		{"1", 0, 0, 0, 0, 0, 0},
		{"2", 0, 0, 0, 180, 20015.086, 0},
		{"3", 0, 0, 90, 0, 10007.543, 0},
		{"4", 0, 0, 90, 180, 10007.543, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := distance(tt.lat1, tt.lon1, tt.lat2, tt.lon2)
			if got < tt.want-tt.want1 || got > tt.want+tt.want1 {
				t.Errorf("distance() = %v, want %v ± %v", got, tt.want, tt.want1)
			}
		})
	}
}

func Test_distanceCompare(t *testing.T) {
	tests := []struct {
		name string
		lat1 float32
		lon1 float32
		lat2 float32
		lon2 float32
		want bool
	}{
		{"1", 0, 0, 0, 0, true},
		{"2", 0, 0, 0, 180, false},
		{"3", 0, 0, 90, 0, false},
		{"4", 0, 0, 90, 180, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distanceCompare(tt.lat1, tt.lon1, tt.lat2, tt.lon2); got != tt.want {
				t.Errorf("distanceCompare() = %v, want %v", got, tt.want)
			}
		})
	}
}
