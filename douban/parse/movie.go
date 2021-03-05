package parse

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//Page struct
type Page struct {
	Page int
	URL  string
}

//Movie struct
type Movie struct {
	Title     string
	Year      int
	Area      string
	Tag       string
	Describes string
	Score     string
	Quote     string
	Ranking   int
}

var (
	baseURL string
	Movies  []Movie
)

func httpParams(req *http.Request) {
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.182 Safari/537.36")
	req.Header.Add("cookie", `viewed="27015617"; bid=ESetovWvBvE; gr_user_id=bad4e190-1608-4e0f-8f59-db8f3c01363b; douban-fav-remind=1; ll="118371"; __utmv=30149280.18438; douban-profile-remind=1; __gads=ID=8c410ff1e4e09e97-2293db1e6fc40041:T=1603977900:RT=1603977900:S=ALNI_MYQJHxghVhhCUm-Bs9A54Zv0J2_MA; __yadk_uid=MMRFiIvIxYoO5Zw1bP1kiwZnqQi1VXyu; _vwo_uuid_v2=DAE4A6D3B901D1FEADC0A04C325A24BE0|e300d5dd2e3f7158148b1e9faadbd593; __utmz=30149280.1612320381.9.9.utmcsr=baidu|utmccn=(organic)|utmcmd=organic; __utmc=30149280; push_noty_num=0; push_doumail_num=0; __utmc=223695111; __utmz=223695111.1613717579.3.3.utmcsr=douban.com|utmccn=(referral)|utmcmd=referral|utmcct=/; _pk_ref.100001.4cf6=%5B%22%22%2C%22%22%2C1613725766%2C%22https%3A%2F%2Fwww.douban.com%2F%22%5D; _pk_ses.100001.4cf6=*; __utma=30149280.1081804443.1599029726.1613723368.1613725766.12; __utmb=30149280.0.10.1613725766; __utma=223695111.412358177.1606448253.1613723368.1613725766.5; __utmb=223695111.0.10.1613725766; dbcl2="184382598:dBtDw27nvVE"; ck=bgvn; _pk_id.100001.4cf6=90dfb0fe92b9cefa.1606448253.5.1613725863.1613723591.`)
}

//Start 开始
func Start(url string) {
	baseURL = url
	client := http.Client{}
	req, err1 := http.NewRequest("GET", url, nil)
	if err1 != nil {
		log.Fatal(err1)
	}
	httpParams(req)
	resp, err2 := client.Do(req)
	if err2 != nil {
		log.Fatal(err2)
	}
	doc, err3 := goquery.NewDocumentFromReader(resp.Body)
	if err3 != nil {
		log.Fatal(err3)
	}
	getAllPage(doc)
	// fmt.Println("开始")
}
func getAllPage(doc *goquery.Document) {
	pages := []Page{}
	pages = append(pages, Page{Page: 1, URL: ""})
	doc.Find(".paginator a").Each(func(i int, s *goquery.Selection) {
		page, _ := strconv.Atoi(s.Text())
		url, _ := s.Attr("href")
		pages = append(pages, Page{Page: page, URL: url})
	})
	pages = pages[:len(pages)-1] //去掉下一页
	pageData(pages)
}
func working(v Page, ch chan string) {
	client := http.Client{}
	req, _ := http.NewRequest("GET", baseURL+v.URL, nil)
	httpParams(req)
	resp, _ := client.Do(req)
	doc, _ := goquery.NewDocumentFromReader(resp.Body)
	doc.Find(".grid_view li").Each(func(i int, s *goquery.Selection) {
		title := s.Find(".hd a .title").Eq(0).Text()
		desc := strings.TrimSpace(s.Find(".bd p").Eq(0).Text())
		descInfo := strings.Split(desc, "\n")
		describes := descInfo[0]
		descOther := strings.Split(strings.TrimSpace(descInfo[1]), "/")
		year, _ := strconv.Atoi((strings.TrimSpace(descOther[0])))
		area := descOther[1]
		tag := descOther[2]
		score := s.Find(".rating_num").Text()
		quote := s.Find(".quote .inq").Text()
		rangking, _ := strconv.Atoi(s.Find(".item .pic em").Text())
		Movies = append(Movies, Movie{Title: title, Year: year, Area: area, Tag: tag, Describes: describes, Score: score, Quote: quote, Ranking: rangking})
	})
	ch <- v.URL

}
func pageData(page []Page) {
	ch := make(chan string)
	for _, v := range page {
		go working(v, ch)
	}
	for i := 0; i < len(page); i++ {
		<-ch
	}
}
