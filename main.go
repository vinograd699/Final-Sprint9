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
	for i := range slice {
		slice[i] = rand.Intn() + 1
	}
	return slice
}

// maximum returns max value in a slice
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

// maxChunks finds max in parallel using goroutines
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	var wg sync.WaitGroup
	chunkSize := (len(data) + CHUNKS - 1) / CHUNKS
	results := make(chan int, CHUNKS)

	// Launch CHUNKS goroutines to process each chunk
	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > len(data) {
			end = len(data)
		}
		if start >= len(data) {
			continue
		}

		wg.Add(1)
		go func(s, e int) {
			defer wg.Done()
			results <- maximum(data[s:e])
		}(start, end)
	}

	// Close results channel when all goroutines are done
	go func() {
		wg.Wait()
		close(results)
	}()

	// Find max among all chunk results
	max := 0
	for v := range results {
		if v > max {
			max = v
		}
	}
	return max
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
