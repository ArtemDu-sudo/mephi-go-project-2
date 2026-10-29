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

	// maxRandomValue — верхняя граница (включительно) для случайных чисел.
	maxRandomValue = 1_000_000
)

// Все числа положительные. Для size <= 0 возвращается пустой слайс.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	data := make([]int, size)
	for i := range data {
		data[i] = rand.Intn(maxRandomValue) + 1
	}

	return data
}

// Для пустого (или nil) слайса возвращается 0: так как все числа положительные,
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	maxValue := data[0]
	for _, value := range data[1:] {
		if value > maxValue {
			maxValue = value
		}
	}

	return maxValue
}

func maxChunks(data []int) int {
	if len(data) < CHUNKS {
		return maximum(data)
	}

	var wg sync.WaitGroup
	chunkSize := len(data) / CHUNKS
	chunkMaxima := make([]int, CHUNKS)

	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize
		// Остаток от деления достаётся последней части, чтобы не потерять элементы.
		if i == CHUNKS-1 {
			end = len(data)
		}

		wg.Add(1)
		go func(index int, chunk []int) {
			defer wg.Done()
			chunkMaxima[index] = maximum(chunk)
		}(i, data[start:end])
	}

	wg.Wait()

	return maximum(chunkMaxima)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	maxValue := maximum(data)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", maxValue, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	maxValue = maxChunks(data)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", maxValue, elapsed)
}
