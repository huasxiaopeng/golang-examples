package main

import "fmt"

/*
  make 创建并初始化好值
 */
func main() {
   //var p*[]int =new([]int)
   //var v[]int =make([]int,100)
	x := []int{1,2,3}
	y := []int{4,5,6}
	x = append(x, y...)
	fmt.Println(x)
}
