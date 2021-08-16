package main

import "fmt"

/**
  嵌套函数的初始化
 */
//Pubs看做父类，Subs 看做子类
type Pubs struct {
   Name string
   age int
}
type Subs struct {
	Pubs
	foodName string
}
//父类构造
func  newPubs(name string,age int)*Pubs  {
	return &Pubs{
		Name: name,
		age: age,
	}
}
//子类构造
func newSubs(foodName string)*Subs  {
	 p:=&Subs{}
	 p.foodName=foodName
	 return p
}

func main() {
	s := newSubs("一哥")
	fmt.Println(s)
}
