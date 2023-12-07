package main

import "fmt"

type person struct {
	name string
	age  int8
	city string
}

//NewPerson 构造函数
func NewPerson(name string, age int8, city string) *person {
	return &person{
		name: name,
		age:  age,
		city: city,
	}
}

func main() {
	var p1 person
	p1.name = "pprof.cn"
	p1.age = 12
	p1.city = "长沙"

	fmt.Printf("p1=%v\n", p1)
	fmt.Printf("p1=%#v\n", p1)

	var user struct {
		Name string
		Age  int
	}
	user.Age = 34
	user.Name = "libing"
	fmt.Printf("%#v\n", user)

	var p2 = new(person)
	p2.name = "测试"
	fmt.Printf("%T\n", p2)
	fmt.Printf("p2=%#v\n", p2)

	p3 := &person{}
	fmt.Printf("%T\n", p3)
	fmt.Printf("p3=%#v\n", p3)

	p5 := person{
		name: "pprof.cn",
		city: "北京",
		age:  18,
	}
	fmt.Printf("p5=%#v\n", p5)

	p6 := &person{
		name: "pprof.cn",
		city: "北京",
		age:  18,
	}
	fmt.Printf("p6=%#v\n", p6)

	p7 := &person{
		city: "北京",
	}
	fmt.Printf("p7=%#v\n", p7)

	p8 := &person{
		"pprof.cn",
		18,
		"北京",
	}
	fmt.Printf("p8=%#v\n", p8)

	type test struct {
		a int8
		b int8
		c int8
		d int8
	}
	n := test{
		1, 2, 3, 4,
	}
	fmt.Printf("n.a %p\n", &n.a)
	fmt.Printf("n.b %p\n", &n.b)
	fmt.Printf("n.c %p\n", &n.c)
	fmt.Printf("n.d %p\n", &n.d)

	type student struct {
		name string
		age  int
	}

	m := make(map[string]*student)
	stus := []student{
		{name: "pprof.cn", age: 18},
		{name: "测试", age: 23},
		{name: "博客", age: 28},
	}

	for _, stu := range stus {
		m[stu.name] = &stu
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}

	p9 := NewPerson("测试", 25, "长沙")
	p9.Dream()
	p9.SetAge(13)
	fmt.Println(p9.age)
	p9.SetAge2(44)
	fmt.Println(p9.age)

	p10 := Person{
		"pprof.cn",
		18,
	}
	fmt.Printf("%#v\n", p10)         //main.Person{string:"pprof.cn", int:18}
	fmt.Println(p10.string, p10.int) //pprof.cn 18

	d1 := &Dog{
		Feet: 4,
		Animal: &Animal{ //注意嵌套的是结构体指针
			name: "乐乐",
		},
	}
	d1.wang() //乐乐会汪汪汪~
	d1.move() //乐乐会动！
}

//Dream Person做梦的方法
func (p person) Dream() {
	fmt.Printf("%s的梦想是学好Go语言！\n", p.name)
}

// SetAge 设置p的年龄
// 使用指针接收者
func (p *person) SetAge(newAge int8) {
	p.age = newAge
}

// SetAge2 设置p的年龄
// 使用值接收者
func (p person) SetAge2(newAge int8) {
	p.age = newAge
}

//Person 结构体Person类型
type Person struct {
	string
	int
}

//Animal 动物
type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会动！\n", a.name)
}

//Dog 狗
type Dog struct {
	Feet    int8
	*Animal //通过嵌套匿名结构体实现继承
}

func (d *Dog) wang() {
	fmt.Printf("%s会汪汪汪~\n", d.name)
}
