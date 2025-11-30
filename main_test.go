package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestGenerateRandomElements checks if generateRandomElements returns correct slice size and values
func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name     string
		size     int
		expected int
	}{
		{"ZeroSize", 0, 0},
		{"NegativeSize", -10, 0},
		{"SmallSize", 5, 5},
		{"LargeSize", 1000, 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			slice := generateRandomElements(tt.size)
			require.NotNil(t, slice)
			assert.Equal(t, tt.expected, len(slice))
			for _, v := range slice {
				assert.Greater(t, v, 0)
			}
		})
	}
}

// TestMaximum checks if maximum function returns correct max value for different inputs
func TestMaximum(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"NilSlice", nil, 0},
		{"EmptySlice", []int{}, 0},
		{"OneElement", []int{99}, 99},
		{"Ascending", []int{10, 20, 30, 40, 50}, 50},
		{"Descending", []int{70, 60, 50, 40, 30}, 70},
		{"Unordered", []int{17, 3, 88, 42, 12}, 88},
		{"AllEqual", []int{111, 111, 111}, 111},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maximum(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
