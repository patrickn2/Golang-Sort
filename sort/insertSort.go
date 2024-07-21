package sort

func InsertSort(list []int) []int {
	for i := 1; i < len(list); i++ {
		currentNum := list[i]
		for z := i - 1; z >= 0; z-- {
			if list[z] < currentNum {
				list[z+1] = currentNum
				break
			}
			list[z+1] = list[z]
			if z == 0 {
				list[z] = currentNum
			}
		}
	}
	return list
}
