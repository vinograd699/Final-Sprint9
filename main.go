package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements creates a slice of random integers
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	slice := make([]int, 0, size)
	rand.Seed(time.Now().UnixNano())
	for i := range slice {
		slice[i] = rand.Intn(1_000_000) + 1
	}
	return slice
}

// maximum returns the max value in a slice
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	max := data[0]
	for _, v := range data[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

// maxChunks finds max in each chunk using goroutines
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	var wg sync.WaitGroup
	maxElements := make([]int, CHUNKS)
	chunkSize := len(data) / CHUNKS
	wg.Add(CHUNKS)

	for i := range CHUNKS {
		start := i * chunkSize
		end := start + chunkSize

		go func(idx, s, e int) {
			defer wg.Done()
			maxElements[idx] = maximum(data[s:e])
		}(i, start, end)
	}
	wg.Wait()
	return maximum(maxElements)
}

func main() {
	fmt.Printf("Generating %d integers\n", SIZE)
	elements := generateRandomElements(SIZE)

	fmt.Println("Finding max in single thread")
	start := time.Now()
	max := maximum(elements)
	elapsed := time.Since(start).Milliseconds()
	fmt.Printf("Max value: %d\nSearch time: %d ms\n", max, elapsed)

	fmt.Printf("Finding max using %d goroutines\n", CHUNKS)
	start = time.Now()
	max = maxChunks(elements)
	elapsed = time.Since(start).Milliseconds()
	fmt.Printf("Max value: %d\nSearch time: %d ms\n", max, elapsed)
}
