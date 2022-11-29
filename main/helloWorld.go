package main

import (
	"TestDemo/dd"
	"fmt"
	"math/rand"
)

func baseDemo() {
	part1()
	part2()
	part3()
	part4()
	part5()
	part6()
	part7()
	part8()
	part8Function()
	part8Function2()
	part9()
	part10()
	part11()
	part12()
	forNote()

}
func forNote() {
	fmt.Println(">>>>>>>>>>>>>>>>>>part 12 note >>>>>>>>>>>>>>>>>>>>>>>>")
	var buffer [256]byte
	slice := buffer[10:20]
	for i := 0; i < len(slice); i++ {
		slice[i] = byte(i)
	}
	fmt.Println("before", slice)
	fmt.Println("buffer index 10: ", &buffer[10])
	fmt.Printf("%p\n", slice)
	AddOneToEachElement(slice)
	fmt.Println("after", slice)
	str1 := "值传递测试str"
	// 一下方法，包括方法内地址打印，验证方法参数传递为值拷贝
	fmt.Println("方法入参为值传递： ")
	fmt.Println("after func str1 value: ", str1)
	fmt.Println(&str1)
	changeStr(str1)
	// 测试指针及值传递type 方法差异
	fmt.Println("============ Path type ============")
	pathName := dd.Path("/usr/bin/tso")
	fmt.Printf("pathName origin: %p\n", pathName)
	pathName.ToUpper()
	fmt.Printf("pathName: %s, %p\n", pathName, pathName)
	fmt.Println("============ Path type ============")
	pathName.TruncateAtFinalSlash()
	fmt.Printf("pathName truncate %s, %p\n", pathName, pathName)
	fmt.Println("============ Path type ============")
	pathName.TruncateAtFinalSlashValue()
	fmt.Printf("pathName value truncate %s, %p\n", pathName, pathName)
	fmt.Println("============ Path type ============")
	pathName.TruncateAtFinalSlash()
	fmt.Printf("pathName truncate2 %s, %p\n", pathName, pathName)
	fmt.Println("============ slice cap ============")
	var iBuffer [10]int
	bufferSlice := iBuffer[0:9]
	fmt.Println(len(bufferSlice))
	fmt.Println(cap(bufferSlice))
	fmt.Printf("iBuffer %p\n", &iBuffer[0])
	fmt.Printf("bufferSlice %p\n", bufferSlice)
	// 扩容，此处是给切片扩容
	bufferSliceExt1 := extend(bufferSlice, 10)
	fmt.Println("extend1:", bufferSlice)
	fmt.Println("bufferSliceExt1: ", bufferSliceExt1)
	fmt.Printf("bufferSliceExt1 p:%p. len: %d, cap: %d\n", bufferSliceExt1, len(bufferSliceExt1), cap(bufferSliceExt1))
	// 再次扩容，尝试将长度扩充超过容量，也就是超过原数组length
	//bufferSliceExt2 := extend(bufferSliceExt1, 20)
	//fmt.Printf("bufferSliceExt2 p:%p. len: %d, cap: %d\n", bufferSliceExt2, len(bufferSliceExt2), cap(bufferSliceExt2))
	// extend 函数内是对原slice再次切片，不能超过slice本身数组长度，此处代码运行失败
	fmt.Println("============ slice make ============")
	makeSlice := make([]int, 8, 10)
	for i := range makeSlice {
		makeSlice[i] = i
	}
	fmt.Printf("makeSlice: %v, %p\n", makeSlice, makeSlice)
	newMakeSlice := append(makeSlice, 999)
	fmt.Printf("newMakeSlice: %v, %p\n", newMakeSlice, newMakeSlice)
	newMakeSlice2 := append(newMakeSlice, 888)
	fmt.Printf("newMakeSlice2: %v, %p\n", newMakeSlice2, newMakeSlice2)
	newMakeSlice3 := append(newMakeSlice2, 777)
	fmt.Printf("newMakeSlice3: %v, %p\n", newMakeSlice3, newMakeSlice3)
	fmt.Println("============ slice copy ============")
	fmt.Printf("makeSlice: %v, %p\n", makeSlice, makeSlice)
	// 扩容一个位置，注意此处扩容，并未超出cap，也就是源数组length
	makeSlice = makeSlice[0 : len(makeSlice)+1]
	copyNums := copy(makeSlice[1:], makeSlice[0:])
	fmt.Printf("makeSlice: %v, copyNums:%d, %p\n", makeSlice, copyNums, makeSlice)
	fmt.Println("============ slice append ============")
	fmt.Printf("makeSlice cap: %d, len: %d\n", cap(makeSlice), len(makeSlice))
}

func part12() {
	fmt.Println(">>>>>>>>>>>>>>>>>>part 12 interface >>>>>>>>>>>>>>>>>>>>>>>>")
	cat := dd.Cat{}
	cat.SetColor("white")
	dog := dd.Dog{}
	dog.SetColor("yellow")
	animals := []dd.Animal{&cat, &dog}
	for _, v := range animals {
		v.Eating()
		v.Sleeping()
	}
}

func part11() {
	fmt.Println(">>>>>>>>>>>>>>>>>>part 11 type>>>>>>>>>>>>>>>>>>>>>>>>")
	var p1 = dd.Point{X: 3, Y: 4}
	var p2 = dd.Point{}
	type point struct {
		x, y int
	}
	var p3 = point{}
	p3.x = 100
	p3.y = 200

	fmt.Println(p1)
	fmt.Println(p2.GetDis(p1))

	var student = dd.Student{}
	student.SetSchool("学校")
	student.SetName("姓名")
	fmt.Printf("%v\n", student)
	fmt.Println(student.Name)
	fmt.Println(student.School)

	fmt.Println("=====1::=======")
	person := dd.Person{}
	superman := dd.Superman{}
	fmt.Printf("orignal person: %p\n", &person)
	superman.Init(1000, &person)
	person.SetAge(20)
	person.Print()
	superman.Print()
	fmt.Println("=====2::=======")
	person2 := Person2{}
	fmt.Printf("%p\n", &person2)
	superman2 := Superman2{ph: 1000}
	superman2.setAge(10)
	fmt.Println(superman2.age)
	superman2.print()
}

func part10() {
	fmt.Println(">>>>>>>>>>>>>>>>>>part 10 map>>>>>>>>>>>>>>>>>>>>>>>>")
	// 使用 make 函数构建map
	// var map_variable map[key_data_type]value_data_type
	// map_variable = make(map[key_data_type]value_data_type)
	map1 := make(map[string]string)
	fmt.Printf("%p\n", map1)
	map1["dd"] = "tt"
	fmt.Println("map1: ", map1)
	mapValue1 := map1["dd"]
	fmt.Println(&mapValue1)
	mapValue2 := map1["dd2"]
	fmt.Println("mapValue2: ", mapValue2)
	fmt.Println("mapValue2 address: ", &mapValue2)
	fmt.Println("mapValue2 is nil ?: ", mapValue2 == "")
	mapValue2 = "tt2"
	fmt.Println(map1)
	fmt.Println("mapValue2 address: ", &mapValue2)
	fmt.Println("mapValue2: ", mapValue2)
	mapValue3, isPresent := map1["dd3"]
	fmt.Println("isPresent", isPresent)
	fmt.Println("mapValue3:", mapValue3)
	map1["dd3"] = "xx"
	rangeMap("map1", map1)
	a5 := []int{11, 22, 33, 44}
	for index, value := range a5 {
		fmt.Print(index, value, "   ")
	}
	fmt.Println()
	for _, value := range a5 {
		fmt.Print(value, "   ")
	}
	fmt.Println()
}

func part9() {
	fmt.Println(">>>>>>>>>>>>>>>>>>part 9 container>>>>>>>>>>>>>>>>>>>>>>>>")
	// 数组
	// var variable_name [SIZE] variable_type
	// 声明数组
	var arr [10]int64
	arr[0] = 12
	arr[9] = 15
	fmt.Println("arr", arr)
	var arr1 = [2]int{10, 20}
	fmt.Println("arr1: ", arr1)
	arr3 := [4]string{"a", "b"}
	fmt.Println("arr3", arr3)
	arr4 := [4]int{1, 2}
	fmt.Println("arr4: ", arr4)
	for i := 0; i < len(arr4); i++ {
		fmt.Printf("arr4 %d =%d\n", i, arr4[i])
	}
	for index, value := range arr {
		println("index: ", index, "value", value)
	}
	fmt.Println(">>>>>>>>>>>>>>>>>>part 9 slice>>>>>>>>>>>>>>>>>>>>>>>>")
	// 切片 切片没有长度限制
	a1 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	s1 := a1[5:10]
	s1[1] = 100
	fmt.Println("s1", s1)
	fmt.Println("a1:", a1)
	fmt.Printf("%p, %p\n", s1, &a1)
	s1 = append(s1, 100, 200)
	fmt.Printf("append s1: %v, %p\n", s1, &s1)
	fmt.Println("append s1: ", s1, &s1)
	printSlice(s1)
	// make
	// make([]T, length, capacity) T: 切片内元素类型 length: 切片初始长度 capacity: 切片容量，默认和length相同，可不传
	// len(s) 返回切片长度 cap(s) 返回切片容量 append(s, T...) copy(s1, s2) 深拷贝
	s2 := make([]int, 2)
	fmt.Println("s2 len: ", len(s2))
	fmt.Println("s2 capacity: ", cap(s2))
	fmt.Printf("s2: %p\n", s2)
	s3 := append(s2, 10, 20)
	fmt.Printf("s2: %p\n", s2)
	fmt.Printf("扩容后的s2: %p\n", s3)
	fmt.Println("s2 appended: ", s2)
	printSlice(s3, "s3")
}

func part8Function2() {
	fmt.Println(">>>>>>>>>>>>>>>>>>part 8 function2>>>>>>>>>>>>>>>>>>>>>>>>")
	subFunc := getSequence()
	fmt.Println(subFunc(true))
	fmt.Println(subFunc(false))
}

func part8Function() {
	fmt.Println(">>>>>>>>>>>>>>>>>>part 8 function>>>>>>>>>>>>>>>>>>>>>>>>")
	// func function_name([parameter list]) [return_type] { function body}
	var sumValue = sum(10, 20)
	fmt.Println(sumValue)
	swapA := 10
	swapB := 20
	i3, i4 := swap(swapA, swapB)
	fmt.Println(i3, i4)
	fmt.Println("swapAB: ", swapA, swapB)
	swap2(swapA, swapB)
	fmt.Println("swap2", swapA, swapB)
	swap3(&swapA, &swapB)
	fmt.Println("swap3", swapA, swapB)
	fmt.Println("func: ", func(a, b int) int {
		return a + b
	}(10, 20))
}

func part8() {
	fmt.Println(">>>>>>>>>>>>>>>>>>part 8 for>>>>>>>>>>>>>>>>>>>>>>>>")
	// for
	for i := 0; i < 100; i++ {
		fmt.Print("i: ", i)
	}
	fmt.Println()
	fmt.Println("while")
	// while
	i := 0
	for i < 100 {
		fmt.Print("i: ", i)
		i++
	}
	fmt.Println()
	fmt.Println("dead")
	// 死循环
	i = 0
	for {
		fmt.Print("i: ", i)
		if i >= 100 {
			break
		}
		i++
	}
	fmt.Println()

	sum1 := 0
	/*	for i := 0; i < 100; i++ {
			sum1 += 1
		}
		fmt.Println(sum1, &sum1)
		sum2 := &sum1
		i := 0
		for i < 100 {
			*sum2 += 1
			i++
		}
		fmt.Println(sum2, *sum2, sum1, &sum1)*/

	i2 := 0
	sum3 := &sum1
	for {
		*sum3 += 1
		i2++
		if i2 >= 100 {
			break
		}
	}
	fmt.Println(sum3, *sum3)
}

func part7() {
	fmt.Println(">>>>>>>>>>>>>>>>>>part 7 switch>>>>>>>>>>>>>>>>>>>>>>>>")
	var str string
	fmt.Println("text in")
	//_, _ = fmt.Scanf("%s", &str)
	str = "xxx"
	switch str {
	case "xxx":
		fmt.Println("start")
	case "a":
		fmt.Println("text A")
	case "b":
		fmt.Println("text b")
	default:
		fmt.Println("text other: ", str)
	}
}

func part6() {
	fmt.Println(">>>>>>>>>>>>>>>>>>part 6 if>>>>>>>>>>>>>>>>>>>>>>>>")
	aa := rand.Int()
	if aa > 100 {
		fmt.Println("if print: if")
	} else if aa < 100 {
		fmt.Println("if print: else if")
	} else {
		fmt.Println("if print: else")
	}
}

func part5() {
	fmt.Println(">>>>>>>>>>>>>>>>>>part 5>>>>>>>>>>>>>>>>>>>>>>>>")
	aa := 10
	bb := &aa
	*bb = 100
	fmt.Println(aa, *bb)
}

func part4() {
	fmt.Println(">>>>>>>>>>>>>>>>>>part 4>>>>>>>>>>>>>>>>>>>>>>>>")
	const (
		login = iota
		logout
		user    = iota + 1
		account = iota + 3
	)
	fmt.Println(login)
	fmt.Println(logout)
	fmt.Println(user)
	fmt.Println(account)
}

func part3() {
	fmt.Println(">>>>>>>>>>>>>>>>>>part 3>>>>>>>>>>>>>>>>>>>>>>>>")
	// 定义枚举
	const (
		Mon, Tue = iota + 1, iota + 2
		Wed, Thu
		Fri, Sat
		Sun, None
	)
	fmt.Println(Mon)
	fmt.Println(Tue)
	fmt.Println(Wed)
	fmt.Println(Thu)
	fmt.Println(Fri)
	fmt.Println(Sat)
	fmt.Println(Sun)
	fmt.Println(None)
}

func part2() {
	fmt.Println(">>>>>>>>>>>>>>>>>>part 2>>>>>>>>>>>>>>>>>>>>>>>>")
	const pi = 3.14
	r := 3.0
	fmt.Println("area:", r*r*pi)
}

func part1() {
	fmt.Println(">>>>>>>>>>>>>>>>>>part 1>>>>>>>>>>>>>>>>>>>>>>>>")
	fmt.Println(">>>>>>>>>>> hi dd!")
	a, b, c := 100, 3.14, "hello world"
	fmt.Printf("%d\n", a)
	fmt.Printf("%4.2f\n", b)
	fmt.Printf("%s\n", c)
}
