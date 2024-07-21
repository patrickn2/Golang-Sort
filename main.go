package main

import (
	"fmt"

	"github.com/patrickn2/Golang-Sort/sort"
)

func main() {
	list := []int{
		20, 6, 1, 8, -8, 2,
	}
	fmt.Println("Quick Sort", sort.QuickSort(list))
	fmt.Println("Insert Sort", sort.InsertSort(list))
	fmt.Println("Bubble Sort", sort.BubbleSort(list))
}
