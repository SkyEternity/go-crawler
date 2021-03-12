package main

import "fmt"

type user struct {
	Name string
}

func main() {
	u := user{}
	u.Name = "haha"
	fmt.Printf("u的值: %v,   内存地址是: %p \n", u, &u)
	change1(u)
	fmt.Printf("u的值: %v,   内存地址是: %p \n", u, &u)
	change2(&u)
	fmt.Printf("u的值: %v,   内存地址是: %p \n", u, &u)

	u.say()
}
func change1(u user) {
	u.Name = "修改了"
	fmt.Printf("u的值: %v,   内存地址是: %p \n", u, &u)
}

func change2(u *user) {
	u.Name = "这次传的是指针类型"
	fmt.Printf("指针类型u的值: %v,   内存地址是: %p \n", u, &u)
}

func (u user) say() {
	fmt.Println("hello")
}
