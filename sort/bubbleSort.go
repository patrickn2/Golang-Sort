package sort

func BubbleSort(list []int) []int {
	if len(list) < 2 {
		return list
	}
	var swapped bool
	index := 0
	for {
		if list[index] > list[index+1] {
			list[index], list[index+1] = list[index+1], list[index]
			swapped = true
		}
		index++
		if index == len(list)-1 {
			index = 0
			if swapped {
				swapped = false
				continue
			}
			return list
		}
	}
}
