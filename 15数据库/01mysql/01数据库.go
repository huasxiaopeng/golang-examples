package main

import (
	"database/sql"

)

func main() {
	db, err := sql.Open("mysql", "root:123456@/mydata?charset=utf8")
	checkErr(err)

	//插入数据
	sql1 := "INSERT INTO user SET name=?,age=?"
	stmt, err := db.Prepare(sql1)
	checkErr(err)
	_, err = stmt.Exec("zs", "30")
	checkErr(err)
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}