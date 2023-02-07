package algorithm

import "fmt"

func quickSort(arr []int) {
	// arr := []int{2, 5, 2, 1, 3, 7, 4, 5, 5, 6, 9,  0}
	//              0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11
	sort := QuickSortDemo2(arr)
	fmt.Println(sort)
}

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	// 第一个数据
	splitData := arr[0]
	low := make([]int, 0)
	high := make([]int, 0)
	mid := make([]int, 0)
	mid = append(mid, splitData)
	for i := 1; i < len(arr); i++ {
		if arr[i] < splitData {
			low = append(low, arr[i])
		} else if arr[i] > splitData {
			high = append(high, arr[i])
		} else {
			mid = append(mid, arr[i])
		}
	}
	low = QuickSort(low)
	high = QuickSort(high)
	myArr := append(append(low, mid...), high...)
	return myArr
}

func QuickSortDemo2(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	splitData := arr[0]
	low := make([]int, 0, 0)
	mid := make([]int, 0, 0)
	high := make([]int, 0, 0)
	mid = append(mid, splitData)
	for i := 1; i < len(arr); i++ {
		if arr[i] > splitData {
			high = append(high, arr[i])
		} else if arr[i] == splitData {
			mid = append(mid, arr[i])
		} else {
			low = append(low, arr[i])
		}
	}
	low = QuickSortDemo2(low)
	high = QuickSortDemo2(high)
	return append(append(low, mid...), high...)
}

func QuickSortDemo(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	point := arr[0]
	low := make([]int, 0)
	high := make([]int, 0)
	mid := make([]int, 0)
	mid = append(mid, point)
	for i := 1; i < len(arr); i++ {
		if arr[i] > point {
			high = append(high, arr[i])
		} else if arr[i] < point {
			low = append(low, arr[i])
		} else {
			mid = append(mid, arr[i])
		}
	}
	// 开始通过递归分别再对各个分区进行排序
	low = QuickSortDemo(low)
	high = QuickSortDemo(high)
	// 开始收集每次递归结果
	result := append(append(low, mid...), high...)
	return result
}
