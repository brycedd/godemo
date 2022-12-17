package algorithm

import (
	"fmt"
)

func Algorithm() {
	arr := []int{2, 5, 2, 1, 3, 7, 4, 5, 5, 6, 9, 0}
	fmt.Println("before:", arr)
	var result []int
	//quickSort(arr)
	//result = BubbleSort(arr)
	result = InsertionSort(arr)
	fmt.Println("after :", result)
}
