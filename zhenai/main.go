package main

import (
	"fmt"
	"go-creeper/zhenai/db"
	"go-creeper/zhenai/model"
	"go-creeper/zhenai/parse"
)

func add(person []model.Person) {
	for _, v := range person {
		sqlStr := "insert into zhenai_user(name,city,salary,education,age,height,introduce,cover) values (?,?,?,?,?,?,?,?)"
		_, err := db.DB.Exec(sqlStr, v.Name, v.City, v.Salary, v.Education, v.Age, v.Height, v.Introduce, v.Cover)
		if err != nil {
			fmt.Printf("insert failed, err:%v\n", err)
			return
		}
	}
	fmt.Println("数据存储完毕")
}

func main() {
	parse.Start()
	add(parse.Persons)
}
