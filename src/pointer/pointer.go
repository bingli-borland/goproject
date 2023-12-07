package main

import "fmt"

func modify1(x int) {
	x = 100
}

func modify2(x *int) {
	*x = 100
}

func main() {
	// a := 10
	// b := &a
	// fmt.Printf("a:%d ptr:%p\n", a, &a)
	// fmt.Printf("b:%p type:%T\n", b, b)
	// fmt.Println(&b)

	// c := *b
	// fmt.Printf("type of c:%T\n", c)
	// fmt.Printf("value of c:%v\n", c)

	// modify1(a)
	// fmt.Println(a) // 10
	// modify2(&a)
	// fmt.Println(a) // 100

	// var p *string
	// fmt.Println(p)
	// fmt.Printf("p的值是%s/n", p)
	// if p != nil {
	// 	fmt.Println("非空")
	// } else {
	// 	fmt.Println("空值")
	// }

	var a *int
	a = new(int)
	*a = 100
	fmt.Println(*a)

	var b map[string]int
	b = make(map[string]int, 10)
	b["测试"] = 100
	fmt.Println(b)

}
