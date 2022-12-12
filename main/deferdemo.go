package main

import (
	"fmt"
	"io"
	"os"
)

func DeferDemo() {
	fmt.Println(test())
	fmt.Println(testDefer2())
	fmt.Println(*testDefer3())
	panicTest()
}

func test() (i int) {
	defer func() {
		i++
	}()
	return 1
}

// 执行testDefer2时，未指定返回值，会生成临时变量作为返回值
// return时为返回值临时变量赋值，此时和i已经没有关系了
func testDefer2() int {
	i := 1
	defer func() {
		i++
	}()
	return i
}

func testDefer3() *int {
	i := 1
	defer func() {
		i++
	}()
	return &i
}

// CopyFile 文件操作
func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer func(src *os.File) {
		err := src.Close()
		if err != nil {

		}
	}(src)

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer func(dst *os.File) {
		err := dst.Close()
		if err != nil {

		}
	}(dst)
	return io.Copy(dst, src)
}

// 捕获异常
func panicTest() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	panic("it is panic")
}
