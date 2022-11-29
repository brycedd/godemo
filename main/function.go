package main

import (
	"fmt"
)

func printSlice(s []int, name ...string) {
	if nil == name || len(name) == 0 {
		fmt.Printf("name = %s len = %d, cap = %d, s =%v\n", name, len(s), cap(s), s)
	} else {
		fmt.Printf("name = %s len = %d, address = %p, s =%v\n", name, len(s), s, s)
	}
}

func rangeMap(mapName string, m map[string]string) {
	if nil == m {
		fmt.Printf("map is nil")
	}
	for k, v := range m {
		fmt.Println(mapName, "key:", k, "value:", v)
	}
}

func getSequence() func(isAdd bool) int {
	i := 0
	return func(isAdd bool) int {
		if isAdd {
			i++
		} else {
			i--
		}
		return i
	}
}

func sum(a, b int) int {
	return a + b
}

func swap(a, b int) (int, int) {
	return b, a
}

func swap2(a, b int) {
	temp := a
	a = b
	b = temp
}

func swap3(a, b *int) {
	*a = *a + *b
	*b = *a - *b
	*a = *a - *b
}

type Person2 struct {
	name string
	age  int
}

func (p *Person2) setAge(age int) {
	p.age = age
}

type Superman2 struct {
	ph int
	Person2
}

func (s Superman2) print() {
	fmt.Printf("persion: %p, superman: %v\n", &s.Person2, s)
}
func AddOneToEachElement(slice []byte) {
	fmt.Printf("%p\n", slice)
	for i := range slice {
		slice[i]++
	}
}

func changeStr(str string) {
	println(&str)
	str = str + "ddd"
}

func extend(s []int, element int) []int {
	i := len(s)
	s = s[0 : i+1]
	s[i] = element
	return s
}

func AppendDD(slice []int, element int) []int {
	l := len(slice)
	if l == cap(slice) {
		// 切片已满，扩容； 为防止s length为0， 扩容为len*2 + 1
		newSlice := make([]int, len(slice), 2*len(slice)+1)
		copy(newSlice, slice)
		slice = newSlice
	}
	// 将slice扩容
	slice = slice[:l+1]
	slice[l] = element
	return slice
}

func Append(slice []int, elements ...int) []int {
	n := len(slice)
	total := n + len(elements)
	if total > cap(slice) {
		// 重新分配为原cap 1.5 倍
		newSize := total*3/2 + 1
		newSlice := make([]int, total, newSize)
		copy(newSlice, slice)
		slice = newSlice
	}
	// 切片扩容
	slice = slice[:total]
	copy(slice[n:], elements)
	return slice
}
