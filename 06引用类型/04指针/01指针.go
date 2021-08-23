package main

import (
	"fmt"
)

func main() {
	//a:=42
	//b:=a
	//fmt.Println(a,b)
	//a=24
	//fmt.Println(a,b)
	var a int =34
	var b *int =&a
	fmt.Println(a,b)
	fmt.Println(&a,b)
	fmt.Println(&a,*b)
	a=22
	fmt.Println(a,*b)

    aa:=[3]int {1,2,3}
    bb:=&aa[0]
    cc:=&aa[1]
    fmt.Println("bb:",bb," cc:",cc)
    fmt.Println("bb:",*bb," cc:",*cc)
}
