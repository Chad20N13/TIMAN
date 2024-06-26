package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("0626.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		return
	}
	fmt.Println("字节数据：", data)
	fmt.Println("字符串数据：", string(data))
	fmt.Println("所获取字节的长度：", count)

	slice001 := make([]int, 10, 100)
	fmt.Println(slice001)
	fmt.Println(len(slice001))
	fmt.Println(cap(slice001))
	slice001[0] = 10
	slice001[1] = 11
	fmt.Println(slice001)
	fmt.Println(len(slice001))
	fmt.Println(cap(slice001))

	slice002 := []int{1, 2, 3, 4}
	fmt.Println(slice002)
	fmt.Println(len(slice002))
	fmt.Println(cap(slice002))

	slice003 := []byte("abc")
	chars := []byte("cedf")
	fmt.Println(slice003)
	fmt.Println(len(slice003))
	fmt.Println(cap(slice003))

	slice003 = append(chars, slice003...)
	fmt.Println(slice003)

}

// 字节数据： [116 104 105 115 32 105 115 32 97 32 116 101 115 116 32 48 54 50 54 46 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
// 字符串数据： this is a test 0626.
// 所获取字节的长度： 20

// func main() {
// 	var conffile = "0626.txt"
// 	fmt.Println("Hello, world.")

// 	//func ReadFile(name string) ([]byte, error) {}
// 	content, err := os.ReadFile(conffile)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(string(content))

// }
