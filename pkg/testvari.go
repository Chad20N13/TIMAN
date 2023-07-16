package abcd

import (

    "fmt"
)

func Dddd() {
            a := "chad001" //define a;
            fmt.Println("define var a: "+a)
            fmt.Println(&a)
            p := &a   //pointer p refer to address of a;
            fmt.Println("pointer p refer to address of a: ",&p)
            fmt.Println(&a)
            fmt.Println("print value of pointer p:" ,*p)
}