package parse

import (
	"go-creeper/zhenai/model"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var Persons = []model.Person{}

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

// 拿到了所有城市的链接
// 爬取每个城市里面的所有数据     先爬取那个城市后爬取那个不确定
// 爬取每页的数据              每个页面的数据不确定
//
func requestPage(city, url string, ch chan string, chc chan string) {
	resp, err := requestFn(url)
	if err != nil {
		log.Fatal(err)
		return
	}
	doc, _ := goquery.NewDocumentFromReader(resp.Body)
	doc.Find(".g-list .list-item").Each(func(i int, s *goquery.Selection) {
		name := s.Find(".content tbody").Eq(0).Find("a").Text()
		imports := s.Find(".content tbody tr").Eq(2).Find("td").Eq(1).Text()
		importsArr := strings.Split(imports, "：")
		var salary string
		var education string
		if strings.Index(importsArr[0], "学") != -1 {
			salary = ""
			education = importsArr[1]
		} else {
			salary = importsArr[1]
			education = ""
		}
		heights := s.Find(".content tbody tr").Eq(3).Find("td").Eq(1).Text()
		heightArr := strings.Split(heights, "：")
		height, _ := strconv.Atoi(heightArr[1])

		ages := s.Find(".content tbody tr").Eq(2).Find("td").Eq(0).Text()
		ageArr := strings.Split(ages, "：")
		age, _ := strconv.Atoi(ageArr[1])
		introduce := s.Find(".content .introduce").Text()
		cover, _ := s.Find(".photo img").Attr("src")
		Persons = append(Persons, model.Person{
			Name:      name,
			City:      city,
			Salary:    salary,
			Education: education,
			Age:       age,
			Height:    height,
			Introduce: introduce,
			Cover:     cover,
		})
	})
	ch <- url
	chc <- url

}

func cityPage(city, url string, chc chan string) {
	//每一个城市最多6页
	var ch = make(chan string)
	for i := 0; i < 6; i++ {
		go requestPage(city, url+"/"+strconv.Itoa(i+1), ch, chc)

	}
	for i := 0; i < 6; i++ {
		<-ch
	}
}

func cityData(citys []model.City) {
	var chc = make(chan string)
	for _, v := range citys {
		go cityPage(v.Name, v.URL, chc)
	}
	for i := 0; i < len(citys); i++ {
		<-chc
	}

}

//3. 获取每一个城市下的所有数据
