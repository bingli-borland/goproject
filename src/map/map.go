package main

import "fmt"

func main() {
	scoremap := make(map[string]int, 8)
	scoremap["小米"] = 90
	scoremap["小二"] = 100
	fmt.Println(scoremap)
	fmt.Println(scoremap["小米"])
	fmt.Printf("type of a:%T\n", scoremap)

	userInfo := map[string]string{
		"username": "pprof.cn",
		"password": "123456",
	}
	fmt.Println(userInfo)

	value, ok := scoremap["小二"]
	fmt.Println(ok, value)

	for k, v := range scoremap {
		fmt.Println(k, v)
	}
}
