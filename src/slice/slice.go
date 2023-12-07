package main

import "fmt"

func main() {
	var s1 []int
	if s1 == nil {
		fmt.Println("是空")
	} else {
		fmt.Println("不是空")
	}
	s2 := []int{}
	var s3 []int = make([]int, 0)
	fmt.Println(s1, s2, s3)
	var s4 []int = make([]int, 0, 0)
	fmt.Println(s4)
	s5 := []int{1, 2, 3}
	fmt.Println(s5)
	arr := [5]int{1, 2, 3, 4, 5}
	var s6 []int
	s6 = arr[1:4]
	fmt.Println(s6)
	s7 := make([]int, 5, 10)
	s7[5] = 1
	fmt.Println(s7)
}
