package main

import (
	"fmt"
	"go-creeper/douban/model"
	"go-creeper/douban/parse"
)

var (
	baseURL = "https://movie.douban.com/top250"
)

//将所有的数据存到数据库中
func add(movies []parse.Movie) {
	for _, v := range movies {
		sqlStr := "insert into douban_movie(title, year, area, tag,describes,score,quote,ranking) values (?,?,?,?,?,?,?,?)"
		_, err := model.DB.Exec(sqlStr, v.Title, v.Year, v.Area, v.Tag, v.Describes, v.Score, v.Quote, v.Ranking)
		if err != nil {
			fmt.Printf("insert failed, err:%v\n", err)
			return
		}
	}
	fmt.Println("数据存储完毕")
	defer model.DB.Close()
}

func main() {
	parse.Start(baseURL)
	add(parse.Movies)
}
