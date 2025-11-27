package main

import (
	"testing"
)

// TestGenerateRandomElements checks edge cases for maximum function
func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"nilSlice", []int{}, 0},               // Empty slice → 0
		{"oneElem", []int{1}, 1},               // One element
		{"equalElem", []int{1, 1, 1, 1}, 1},    // All values equal
		{"maxStart", []int{100, 2, 3, 4}, 100}, // Max at start
		{"maxEnd", []int{1, 2, 3, 999}, 999},   // Max at end
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.data); got != tt.want {
				t.Errorf("maximum() = %d, want %d", got, tt.want)
			}
		})
	}
}

// TestMaximum verifies maximum function with various inputs
func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"Empty slice", []int{}, 0},                    // Empty slice
		{"Single element", []int{42}, 42},              // One value
		{"Multiple elements", []int{1, 3, 2, 5, 4}, 5}, // Mixed values
		{"All equal", []int{7, 7, 7}, 7},               // Identical items
		{"Mixed numbers", []int{1, 0, 10, 5}, 10},      // Unordered
		{"Nil slice", nil, 0},                          // Nil input
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximum(tt.data)
			if got != tt.want {
				t.Errorf("maximum(%v) = %d, want %d", tt.data, got, tt.want)
			}
		})
	}
}
