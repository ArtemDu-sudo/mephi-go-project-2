package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{
		{name: "negative size", size: -5},
		{name: "zero size", size: 0},
		{name: "one element", size: 1},
		{name: "small size", size: 10},
		{name: "large size", size: 10_000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateRandomElements(tt.size)

			wantLen := tt.size
			if wantLen < 0 {
				wantLen = 0
			}
			if got == nil {
				t.Fatalf("generateRandomElements(%d) returned nil slice", tt.size)
			}
			if len(got) != wantLen {
				t.Fatalf("generateRandomElements(%d) len = %d, want %d", tt.size, len(got), wantLen)
			}
			for i, value := range got {
				if value <= 0 {
					t.Fatalf("element %d = %d, want positive number", i, value)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "nil slice", data: nil, want: 0},
		{name: "empty slice", data: []int{}, want: 0},
		{name: "one element", data: []int{42}, want: 42},
		{name: "two elements", data: []int{3, 7}, want: 7},
		{name: "max at start", data: []int{9, 1, 2, 3}, want: 9},
		{name: "max in middle", data: []int{1, 2, 10, 3, 4}, want: 10},
		{name: "max at end", data: []int{1, 2, 3, 11}, want: 11},
		{name: "all equal", data: []int{5, 5, 5, 5}, want: 5},
		{name: "descending", data: []int{8, 7, 6, 5}, want: 8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.data); got != tt.want {
				t.Errorf("maximum(%v) = %d, want %d", tt.data, got, tt.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	// Слайс длиной 20 не делится на CHUNKS нацело: остаток уходит в последнюю часть.
	withMaxAt := func(length, position, value int) []int {
		data := make([]int, length)
		for i := range data {
			data[i] = 1
		}
		data[position] = value
		return data
	}

	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "nil slice", data: nil, want: 0},
		{name: "empty slice", data: []int{}, want: 0},
		{name: "one element", data: []int{42}, want: 42},
		{name: "fewer elements than chunks", data: []int{1, 5, 3}, want: 5},
		{name: "exactly CHUNKS elements", data: []int{1, 2, 3, 4, 5, 6, 7, 8}, want: 8},
		{name: "divisible, max at start", data: withMaxAt(16, 0, 99), want: 99},
		{name: "divisible, max at end", data: withMaxAt(16, 15, 99), want: 99},
		{name: "not divisible, max in remainder", data: withMaxAt(20, 19, 99), want: 99},
		{name: "not divisible, max in middle", data: withMaxAt(20, 9, 99), want: 99},
		{name: "all equal", data: []int{7, 7, 7, 7, 7, 7, 7, 7, 7, 7}, want: 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxChunks(tt.data); got != tt.want {
				t.Errorf("maxChunks(%v) = %d, want %d", tt.data, got, tt.want)
			}
		})
	}
}

func TestMaxChunksMatchesMaximum(t *testing.T) {
	for _, size := range []int{0, 1, 7, 8, 9, 100, 1_001, 100_003} {
		data := generateRandomElements(size)
		if got, want := maxChunks(data), maximum(data); got != want {
			t.Errorf("size %d: maxChunks = %d, maximum = %d", size, got, want)
		}
	}
}
