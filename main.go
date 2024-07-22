package main

import (
	"fmt"
	"time"

	"math/rand"

	"github.com/patrickn2/Golang-Sort/sort"
)

func main() {
	// list := []int{
	// 	20, 6, 1, 8, -8, 2, 7, 34, 647, 86, 12, 43, 75, 99, 12, 43, 655, 432, 23623, 623, 267, 323, 7457, 2346, 45754, 2347, 8989, 2346, 24, 636,
	// }

	const sliceSize = 100000
	list := make([]int, sliceSize)

	for i := 0; i < sliceSize; i++ {
		list[i] = rand.Intn(100000) // Adjust the upper bound as needed
	}

	startTime := time.Now()
	sort.MergeSort(list)
	// fmt.Println("Merge Result ", result)
	fmt.Printf("Merge Sort Time: %v\n", time.Since(startTime))
	startTime = time.Now()
	sort.QuickSort(list)
	// fmt.Println("Quick Result ", result)
	fmt.Printf("Quick Sort Time: %v\n", time.Since(startTime))
	startTime = time.Now()
	sort.InsertSort(list)
	// fmt.Println("Insert Result ", result)
	fmt.Printf("Insert Sort Time: %v\n", time.Since(startTime))
	startTime = time.Now()
	sort.BubbleSort(list)
	// fmt.Println("Bubble Result ", result)
	fmt.Printf("Bubble Sort Time: %v\n", time.Since(startTime))
}
