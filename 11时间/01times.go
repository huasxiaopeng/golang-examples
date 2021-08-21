package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("%T\n",now)
	fmt.Printf("%t\n",now)
	fmt.Println(now)

	//自定义时间
	customTime := time.Date(2008, 7, 15, 13, 30,0,0, time.Local)
	fmt.Println(customTime)

	//时间解析
	format := now.Format("2006年1月2日 15:04:05")// 必须传入出来的日期，多么的自恋
	fmt.Println("舒服的格式化：",format)
	stringTime := "2019-01-01 15:03:01"
	parse, _ := time.Parse("2006-01-02 15:04:05", stringTime)
	fmt.Println("相应的解析",parse)


	nowTime := time.Now()

	year, month, day := nowTime.Date()
	fmt.Println(year, month, day)		// 2019 November 01

	hour, min, sec := nowTime.Clock()
	fmt.Println(hour, min, sec)

	fmt.Println(nowTime.Year())
	fmt.Println(nowTime.Month())
	fmt.Println(nowTime.Hour())

	fmt.Println(nowTime.YearDay())

	//时间戳
	fmt.Println(nowTime.Unix())
	//时间睡眠
	time.Sleep(time.Second * 3)
}
