package sort

func QuickSort(list []int) []int {
	pivot := list[len(list)-1]
	var left, right []int
	for i := 0; i < len(list)-1; i++ {
		if list[i] > pivot {
			right = append(right, list[i])
			continue
		}
		left = append(left, list[i])
	}

	if len(left) > 1 {
		left = QuickSort(left)
	}
	if len(right) > 1 {
		right = QuickSort(right)
	}
	list = append(left, pivot)
	list = append(list, right...)
	return list
}
