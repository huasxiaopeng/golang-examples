package main

import "fmt"

func main() {
/**
	for 的三种定义方式
	// Like a C for
  for init; condition; post { }

  // Like a C while
  for condition { }

  // Like a C for(;;)
  for { }
 */
	sum:=0
	for i:=0;i<10;i++{
	  sum+=i;

	}
	fmt.Println(sum)
//If you're looping over an array, slice, string, or map, or reading from a channel,
//a range clause can manage the loop.
/**
  for key, value := range oldMap {
      newMap[key] = value
  }
 */
/**
If you only need the second item in the range (the value),
use the blank identifier,
an underscore, to discard the first:
 */
  var arr=[]int{1,2,3,4,5,6,8}
	for i, value := range arr {
		fmt.Println("i:",i,"  value:",value)
	}
	//有变量不需要 ，直接用_替代
	for _, value := range arr {
		fmt.Println(" value:",value)
	}
	var a=[6]int{}
	// Reverse a
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
