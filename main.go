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

	slice := make([]int, size)
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range slice {
		slice[i] = rng.Int() + 1 // Ensure positive values
	}
	return slice
}

// maximum returns the max value in a slice
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	max := data[0]
	for i := 1; i < len(data); i++ {
		if data[i] > max {
			max = data[i]
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
	var mu sync.Mutex // Protects access to maxElements
	maxElements := make([]int, CHUNKS)
	chunkSize := len(data) / CHUNKS
	wg.Add(CHUNKS)

	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize

		// Handle tail in the last chunk
		if i == CHUNKS-1 {
			end = len(data)
		}

		chunk := data[start:end]
		go func(idx int, c []int) {
			defer wg.Done()
			m := maximum(c)
			mu.Lock()
			maxElements[idx] = m
			mu.Unlock()
		}(i, chunk)
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
	maxParallel := maxChunks(elements)
	elapsed = time.Since(start).Milliseconds()
	fmt.Printf("Max value: %d\nSearch time: %d ms\n", maxParallel, elapsed)

	if max != maxParallel {
		fmt.Printf("WARNING: Results differ! Single: %d, Parallel: %d\n", max, maxParallel)
	}
}
