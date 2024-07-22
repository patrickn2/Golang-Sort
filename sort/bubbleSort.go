package sort

func BubbleSort(list []int) (result []int) {
	result = append(result, list...)
	if len(result) < 2 {
		return result
	}
	var swapped bool
	index := 0
	for {
		if result[index] > result[index+1] {
			result[index], result[index+1] = result[index+1], result[index]
			swapped = true
		}
		index++
		if index == len(result)-1 {
			index = 0
			if swapped {
				swapped = false
				continue
			}
			return result
		}
	}
}
