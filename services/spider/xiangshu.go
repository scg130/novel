package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func main() {
	xiangshu("剑道第一仙")
}

func xiangshu(novelName string) {
	url := "https://www.xbiquge.la/modules/article/waps.php"
	req, err := http.NewRequest("GET", url, strings.NewReader("searchkey="+novelName))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Connection", "close")
	req.Header.Add("Content-type", "application/x-www-form-urlencoded")

	cli := &http.Client{}
	res, err := cli.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode != 200 {
		log.Fatal("request fail")
	}
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".grid").Find("tr").Each(func(i int, s *goquery.Selection) {
		detailUrl, exist := s.Find(".even").Find("a").Attr("href")
		if !exist {
			return
		}
		name := s.Find(".even").Find("a").Text()
		if name != novelName {
			return
		}
		xiangshuNovel(detailUrl)
	})
}

func xiangshuNovel(url string) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode != 200 {
		log.Fatal("request fail")
	}
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	img, _ := doc.Find("#fmimg").Find("img").Attr("src")
	title, _ := doc.Find("#info").Find("h1").Html()
	author, _ := doc.Find("#info").Find("p").First().Html()
	intro, _ := doc.Find("#intro").Find("p").Last().Html()
	sql := "select * from novel.novel where name = ?"
	result := make(map[string]interface{})
	exist, err := x.Table("novel.novel").Where("name = ?", title).Get(&result)
	if err != nil {
		log.Fatal(err)
	}
	var novelId int64
	if exist {
		novelId = result["id"].(int64)
	} else {
		sql = "insert into novel(name,author,img,intro,cate_id) values (?, ?, ?, ?, 5)"
		insertRes, err := x.Exec(sql, title, author, img, intro)
		if err == nil && insertRes != nil {
			novelId, _ = insertRes.LastInsertId()
		}
	}

	doc.Find("#list").Find("dd").Each(func(i int, s *goquery.Selection) {
		// if i < 2640 {
		// 	return
		// }
		chapterTitle := s.Find("a").Text()
		href, _ := s.Find("a").Attr("href")
		href = "https://www.xbiquge.la" + href
		xiangshuChapter(href, chapterTitle, i+1, int(novelId))
	})
}

func xiangshuChapter(url, title string, num, novelId int) {
	rsp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer rsp.Body.Close()
	if rsp.StatusCode != 200 {
		log.Fatal("request fail")
	}

	doc, err := goquery.NewDocumentFromReader(rsp.Body)
	if err != nil {
		log.Fatal(err)
	}
	content, _ := doc.Find("#content").Html()
	sql := "insert into chapter(title,content,novel_id,num,words) values (?, ?, ?, ?, ?)"
	var t string
	if len(strings.Split(title, " ")) == 2 {
		t = strings.Split(title, " ")[1]
	}
	_, err = x.Exec(sql, t, content, novelId, num, len(content))
	if err != nil {
		log.Fatal(err)
	}
	sql = "update novel set chapter_total = chapter_total+1,chapter_current=?,words = words+? where id = ?"
	_, err = x.Exec(sql, num+1, len(content), novelId)
	fmt.Println(err)
}

const (
	user = "root"
	pass = "smd123456"
	db   = "novel"
	port = "3306"
	host = "60.205.191.37"
)

var x *xorm.Engine

func initDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", user, pass, host, port, db)
	var err error
	x, err = xorm.NewEngine("mysql", dsn)
	if err != nil {
		log.Fatalf("mysql connect err:%s", err.Error())
	}
	if err := x.Ping(); err != nil {
		log.Fatalf("mysql ping err:%s", err.Error())
	}
	x.ShowSQL(true)

	x.SetMaxIdleConns(5)

	x.SetMaxOpenConns(5)
}

func init() {
	initDB()
}
