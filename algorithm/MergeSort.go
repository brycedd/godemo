package algorithm

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	p := len(arr) / 2
	left := MergeSort(arr[:p])
	right := MergeSort(arr[p:])
	return Merge(left, right)
}
func Merge(left, right []int) []int {
	i, j, index := 0, 0, 0
	resultArr := make([]int, len(left)+len(right))
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			resultArr[index] = left[i]
			index++
			i++
		} else {
			resultArr[index] = right[j]
			index++
			j++
		}
	}
	for i < len(left) {
		resultArr[index] = left[i]
		index++
		i++
	}
	for j < len(right) {
		resultArr[index] = right[j]
		index++
		j++
	}
	return resultArr
}
