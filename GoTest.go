package main

import "fmt"

type Data struct {
	x int
}

func (self Data) ValueTest() { // 值接受者方法
	fmt.Printf("值接收者: %p\n", &self)
}
func (self *Data) PointerTest() { // 指针接受者方法
	fmt.Printf("指针接受者: %p\n", self)
}

type User struct {
	name string
	age  int
}

func main() {
	ma := make(map[int]*User)
	andes := User{
		name: "andes",
		age:  19,
	}
	ma[1] = &andes
	andes.age = 20

	fmt.Println(ma[1].age)
	fmt.Println(andes.age)
}
