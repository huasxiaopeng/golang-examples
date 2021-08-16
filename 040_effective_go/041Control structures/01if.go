package main

import (
	"fmt"
	"os"
)

func main() {
 //simple
	x:=10
	if x>0{
	//强制使用大括号，
	/**
	It's good style to do so anyway,
		especially when the body contains a
		control statement such as a return or break.
		  保持好的风格
	 */
	}
	/**
	Since if and switch accept an initialization statement,
	it's common to see one used to set up a local variable.

	*/
    if a:=10;a>0{
    	fmt.Println("a的值为：",a)
	}else{
		fmt.Println("a的值为",a)
	}

	open, err := os.Open("name")
	if err != nil {
	  fmt.Println("直接进行数据处理")
	}
	fmt.Println(open)


/**
    当主题存在这些时，会省略执行else
  break, continue, goto, or return—the unnecessary else is omitted.
 */
  if b:=10;b<10{
  	fmt.Println("b")
	  return
  }
	//f, err := os.Open("name")
	//if err != nil {
	//	return err
	//}
	//d, err := f.Stat()
	//if err != nil {
	//	f.Close()
	//	return err
	//}


}
