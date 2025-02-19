package algorithm

import (
	"fmt"
)

func Algorithm() {
	arr := []int{2, 5, 2, 1, 3, 7, 4, 5, 5, 6, 9, 0}
	//arr := []int{3, 2}
	fmt.Println("before:", arr)
	var result []int
	result = quickSort(arr)
	//result = BubbleSort(arr)
	//result = InsertionSort(arr)
	//result = MergeSort(arr)
	//result = SelectionSort(arr)
	//result = HeapSort(arr)
	fmt.Println("after :", result)
}
