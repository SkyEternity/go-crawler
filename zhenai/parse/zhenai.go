package parse

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

//City 城市struct
type City struct {
	Name string
	URL  string
}

func httpParams(req *http.Request) {
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.192 Safari/537.36")
	req.Header.Add("cookie", `sid=69d048d4-e6f2-4858-8c8f-2ec420c54f11; __channelId=900122%2C0; ec=q5JKvSet-1614577173010-9ceed1be893c9-1416558426; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1614578218; _efmdata=Jt2FeBjL1DZpWU3T3h%2FMN4q0bdOkuzQSPGJmIteK9NgPcnNuKrNV%2BRj0roG%2FUPQ3heY1eDG8iQqqIaIZIkFGs2zXCsqissKYbKKdhyghRyA%3D; _exid=KdJTQZhW%2BJSDEuUSiJpRac%2BQjS71Cf%2B2UXvhqgfRWwiaZoEWgdm4xSyL9PLBnUIgmqh02m9o%2FXSpYLog1tfFLg%3D%3D; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1615174515`)
}

//Start 1. 获取所有的城市url
func Start() {
	citys := []City{}
	client := http.Client{}
	req, _ := http.NewRequest("GET", "https://www.zhenai.com/zhenghun", nil)
	httpParams(req)
	resp, _ := client.Do(req)
	doc, _ := goquery.NewDocumentFromReader(resp.Body)
	defer resp.Body.Close()
	doc.Find(".city-list dd").Each(func(i int, s *goquery.Selection) {
		s.Find("a").Each(func(i int, a *goquery.Selection) {
			name := a.Text()
			url, _ := a.Attr("href")
			citys = append(citys, City{Name: name, URL: url})
		})
	})
	fmt.Println(citys)
}
