package sort

func InsertSort(list []int) (result []int) {
	result = append(result, list...)
	for i := 1; i < len(result); i++ {
		currentNum := result[i]
		for z := i - 1; z >= 0; z-- {
			if result[z] < currentNum {
				result[z+1] = currentNum
				break
			}
			result[z+1] = result[z]
			if z == 0 {
				result[z] = currentNum
			}
		}
	}
	return result
}
