package test

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func Readfile() {
	file, err := os.OpenFile("0626.txt", os.O_RDONLY, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		if lineData, err := reader.ReadString('\n'); err != nil {
			if err == io.EOF {
				// 因为是以换行符为分隔符，如果最后一行没有换行符，那么返回 io.EOF 错误时，也是可能读到数据的，因此判断一下是否读到了数据
				if len(lineData) > 0 {
					fmt.Println(lineData)
				}
				break
			}
		} else {
			fmt.Println(strings.TrimRight(lineData, "\n"))
		}
	}
}
