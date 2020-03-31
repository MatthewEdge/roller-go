package math

import (
	"fmt"
	"testing"
)

func TestMin(t *testing.T) {
	tests := []struct {
		a        int
		b        int
		expected int
	}{
		{1, 2, 1},
		{657, 45, 45},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("Min: %d | %d", tt.a, tt.b)
		t.Run(name, func(t *testing.T) {
			result := Min(tt.a, tt.b)
			if result != tt.expected {
				fmt.Printf("Expected %d. Got %d\n", tt.expected, result)
				t.Fail()
			}
		})
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		a   int
		b   int
		out int
	}{
		{1, 2, 2},
		{657, 45, 657},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("Max: %d | %d", tt.a, tt.b)
		t.Run(name, func(t *testing.T) {
			result := Max(tt.a, tt.b)
			if result != tt.out {
				fmt.Printf("Expected %d. Got %d\n", tt.out, result)
				t.Fail()
			}
		})
	}
}
