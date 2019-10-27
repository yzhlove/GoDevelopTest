package main

import (
	"WorkSpace/GoDevelopTest/chat04/obj"
	"fmt"
)

func main() {

	userList := obj.New()

	s := obj.Student{
		Name: "yzh",
		Age:  16,
	}

	fmt.Println(s.ToString())
	s.Show()

	t := obj.Teacher{
		ID:   1,
		Name: "lcm",
		Age:  16,
	}

	fmt.Printf("%T\n", s.ToString)

	userList.Add(&s)
	userList.Add(t)

	userList.To()

}
