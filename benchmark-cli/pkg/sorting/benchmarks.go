package sorting

import (
	"time"
)

type Algorithm int

const (
	BubbleSortAlg Algorithm = iota
	MergeSortAlg
	QuickSortAlg
)

func (a Algorithm) String() string {
	return [...]string{"Bubble Sort", "Merge Sort", "Quick Sort"}[a]
}

type BenchmarkResult struct {
	Algorithm Algorithm
	Time      time.Duration
}

type SortFunc func([]int)

func Benchmark(arr []int) []BenchmarkResult {
	allAlgorithms := []struct {
		name      Algorithm
		algorithm SortFunc
	}{
		{BubbleSortAlg, BubbleSort},
		{MergeSortAlg, MergeSort},
		{QuickSortAlg, QuickSort},
	}

	var results []BenchmarkResult

	for _, algorithm := range allAlgorithms {
		start := time.Now()
		algorithm.algorithm(arr)
		end := time.Now()
		results = append(results, BenchmarkResult{
			Algorithm: algorithm.name,
			Time:      end.Sub(start),
		})
	}

	return results
}
