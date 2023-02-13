package algorithm

func InsertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		preIndex := i - 1
		currentValue := arr[i]
		for preIndex >= 0 && arr[preIndex] > currentValue {
			arr[preIndex], arr[preIndex+1] = currentValue, arr[preIndex]
			preIndex--
		}
	}
	return arr
}

func InsertionSort2(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		preIndex := i - 1
		currentValue := arr[i]
		if preIndex >= 0 && arr[preIndex] > currentValue {
			arr[preIndex], arr[preIndex+1] = currentValue, arr[preIndex]
			preIndex--
		}
	}
	return arr
}
