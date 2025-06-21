package utils

import (
	"fmt"
	"testing"
)

func TestMinutesBetween(t *testing.T) {
	tests := []struct {
		t1, t2 string
		want   int
	}{
		{"12:00", "12:00", 0},
		{"12:00", "12:30", 30},
		{"12:30", "13:00", 30},
		{"12:00", "13:00", 60},
		{"12:00", "12:15", 15},
		{"12:15", "12:00", -15},   // Negative case
		{"23:59", "00:01", -1438}, // Edge case crossing midnight is "wrong" from a human perspective, known limitation
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s to %s", tt.t1, tt.t2), func(t *testing.T) {
			got, err := MinutesBetween(tt.t1, tt.t2)

			if err != nil {
				t.Errorf("MinutesBetween(%q, %q) returned error: %v", tt.t1, tt.t2, err)
			}

			if got != tt.want {
				t.Errorf("MinutesBetween(%q, %q) = %d; want %d", tt.t1, tt.t2, got, tt.want)
			}
		})
	}
}
