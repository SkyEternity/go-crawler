package parse

import (
	"fmt"
	"go-creeper/zhenai/model"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func requestFn(url string) (resp *http.Response, err error) {
	client := http.Client{}
	req, err1 := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.192 Safari/537.36")
	if err1 != nil {
		err = err1
		return
	}
	resp, err = client.Do(req)
	if err != nil {
		return
	}
	return
}

//Start 1. 获取所有的城市url
func Start() {
	resp, err := requestFn("https://www.zhenai.com/zhenghun")
	if err != nil {
		log.Fatal(err)
	}
	citys := []model.City{}
	doc, _ := goquery.NewDocumentFromReader(resp.Body)
	defer resp.Body.Close()
	doc.Find(".city-list dd").Each(func(i int, s *goquery.Selection) {
		s.Find("a").Each(func(i int, a *goquery.Selection) {
			name := a.Text()
			url, _ := a.Attr("href")
			citys = append(citys, model.City{Name: name, URL: url})
		})
	})
	cityData(citys)
}

//2. 获取每个城市下的列表
func cityData(citys []model.City) {
	persons := []model.Person{}
	for _, v := range citys {
		resp, err := requestFn(v.URL)
		if err != nil {
			log.Fatal(err)
			continue
		}
		doc, _ := goquery.NewDocumentFromReader(resp.Body)
		fmt.Println(doc)
		doc.Find(".g-list .list-item").Each(func(i int, s *goquery.Selection) {
			name := s.Find(".content tbody").Eq(0).Find("a").Text()
			city := v.Name
			imports := s.Find(".content tbody tr").Eq(2).Find("td").Eq(1).Text()
			importsArr := strings.Split(imports, "：")
			var salary string
			var education string
			if importsArr[0] == "学历" {
				salary = ""
				education = importsArr[1]
			} else {
				salary = importsArr[1]
				education = ""
			}
			persons = append(persons, model.Person{
				Name:      name,
				City:      city,
				Salary:    salary,
				Education: education,
				Age:       0,
				Height:    0,
				Introduce: "",
				Cover:     "",
			})
		})
		panic(1111)

	}

}
