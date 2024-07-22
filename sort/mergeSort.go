package sort

func MergeSort(list []int) []int {
	if len(list) < 2 {
		return list
	}
	left := list[:len(list)/2]
	right := list[len(list)/2:]

	return merge(MergeSort(left), MergeSort(right))
}

func merge(leftList []int, rightList []int) []int {
	var sortedList []int

	for len(leftList) > 0 && len(rightList) > 0 {
		if leftList[0] <= rightList[0] {
			item := leftList[0]
			leftList = leftList[1:]
			sortedList = append(sortedList, item)
			continue
		}
		item := rightList[0]
		rightList = rightList[1:]
		sortedList = append(sortedList, item)
	}
	sortedList = append(sortedList, leftList...)
	sortedList = append(sortedList, rightList...)
	return sortedList
}
