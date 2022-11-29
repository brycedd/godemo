package main

import "fmt"

func testFunc() {
	str := "xxx"
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
	// print "start"
}
