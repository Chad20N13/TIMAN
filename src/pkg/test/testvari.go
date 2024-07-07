package test

import (
	"fmt"
)

func Dddd() {

	defer fmt.Printf("test defer1")
	defer fmt.Printf("test defer2")
	a := "chad001" //define a;

	fmt.Println("define var a: " + a)

	fmt.Println("memory address of a :", &a)

	p := &a //define pointer P , directing to address of a,  实际上这个p里面存放的东西，就是变量a的内存地址信息

	fmt.Println("print value of pointer p:", *p)

	var ptr *string = &a
	fmt.Println("print value of pointer p:", ptr)
	fmt.Println("print value of a:", *ptr)

}
