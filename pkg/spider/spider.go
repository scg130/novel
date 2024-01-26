package spider

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/axgle/mahonia"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/sirupsen/logrus"
)

const URL = "http://m.b5200.net/modules/article/waps.php?keyword=%s"

var lockMap *sync.Map
var enc mahonia.Decoder

func init() {
	enc = mahonia.NewDecoder("gbk")
	initDB()
	lockMap = &sync.Map{}
}

func Run(pyName, zhName string) error {
	key := fmt.Sprintf("%s:%s", pyName, zhName)
	if _, ok := lockMap.Load(key); ok {
		logrus.Error("task exist!")
		return errors.New("task exist!")
	}
	lockMap.Store(key, 1)
	defer lockMap.Delete(key)
	err := Xiangshu(zhName)
	if err != nil {
		logrus.Error(err)
		// find(zhName)
		return err
	}
	return nil
}

func find(name string) error {
	names := strings.Split(name, "|")
	url := fmt.Sprintf(URL, names[0])
	res, err := http.Get(url)
	if err != nil {
		logrus.Error(err)
		return err
	}
	if res.StatusCode != 200 {
		logrus.Error("request fail 1")
		return errors.New("fail")
	}
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		logrus.Error(err)
		return err
	}
	flag := true
	doc.Find(".cover").Find("p").Each(func(i int, s *goquery.Selection) {
		author := enc.ConvertString(s.Find("a").Last().Text())
		novelName := enc.ConvertString(s.Find(".blue").Text())
		detailUrl, _ := s.Find(".blue").Attr("href")
		if author == names[1] && novelName == names[0] {
			flag = false
			novel(fmt.Sprintf("http://m.b5200.net%s", detailUrl), author)
		}
	})
	if flag {
		return errors.New("not found")
	}
	return nil
}

func novel(url string, author string) error {
	res, err := http.Get(url)
	if err != nil {
		logrus.Error(err)
		return err
	}
	if res.StatusCode != 200 {
		logrus.Error("request fail 1")
		return errors.New("fail")
	}
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		logrus.Error(err)
		return err
	}
	img, _ := doc.Find(".block_img2").Find("img").Attr("src")
	title, _ := doc.Find("#block_txt2").Find("h2").Find("a").Html()
	title = enc.ConvertString(title)
	intro, _ := doc.Find("#info").Html()
	sql := "select * from novel.novel where name = ?"
	result := make(map[string]interface{})
	exist, err := x.Table("novel.novel").Where("name = ?", title).Get(&result)
	if err != nil {
		logrus.Error(err)
		return err
	}
	var novelId, chapterCurent int64
	if exist {
		chapterCurent = result["chapter_current"].(int64)
		novelId = result["id"].(int64)
	} else {
		sql = "insert into novel(name,author,img,intro,cate_id) values (?, ?, ?, ?, 1)"
		insertRes, err := x.Exec(sql, title, author, img, intro)
		if err == nil && insertRes != nil {
			novelId, _ = insertRes.LastInsertId()
		}
	}

	listUrl, _ := doc.Find(".ablum_read").First().Find("span").Last().Find("a").Attr("href")
	url1 := "http://m.b5200.net" + listUrl[:len(listUrl)-1] + "_1/"
	res, err = http.Get(url1)
	if err != nil {
		logrus.Error(err)
		return err
	}
	if res.StatusCode != 200 {
		logrus.Error("request fail 1")
		return errors.New("fail")
	}
	defer res.Body.Close()

	doc, err = goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		logrus.Error(err)
		return err
	}
	page := enc.ConvertString(doc.Find(".page").Last().Text())
	pages, _ := strconv.Atoi(page[strings.Index(page, "/")+1 : strings.Index(page, ")")-3])

	for i := 1; i <= pages; i++ {
		chaptersUrl := fmt.Sprintf("http://m.b5200.net"+listUrl[:len(listUrl)-1]+"_%d/", i)
		chapters(chaptersUrl, chapterCurent, novelId, i)
	}
	return nil
}

func chapters(url string, chapterCurent int64, novelId int64, page int) error {
	res, err := http.Get(url)
	if err != nil {
		logrus.Error(err)
		return err
	}
	if res.StatusCode != 200 {
		logrus.Error("request fail 1")
		return errors.New("fail")
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		logrus.Error(err)
		return err
	}
	doc.Find(".chapter").Find("li").Each(func(i int, s *goquery.Selection) {
		if i+(20*page-1) < int(chapterCurent) {
			return
		}

		href, _ := s.Find("a").Attr("href")
		chapter(fmt.Sprintf("http://m.b5200.net%s", href), enc.ConvertString(s.Text()), i, int(novelId))
	})
	return nil
}

func chapter(url, title string, num, novelId int) {
	fmt.Println(url)
	rsp, err := http.Get(url)
	if err != nil {
		logrus.Error(err)
		return
	}

	defer rsp.Body.Close()
	if rsp.StatusCode != 200 {
		logrus.Error("request fail chapter")
		return
	}

	doc, err := goquery.NewDocumentFromReader(rsp.Body)
	if err != nil {
		logrus.Error(err)
		return
	}
	content, _ := doc.Find(".text").Html()
	content = enc.ConvertString(content)
	sql := "insert into chapter(title,content,novel_id,num,words) values (?, ?, ?, ?, ?)"
	var t string
	if len(strings.Split(title, " ")) == 2 {
		t = strings.Split(title, " ")[1]
	}
	_, err = x.Exec(sql, t, content, novelId, num+1, len(content))
	if err != nil {
		logrus.Error(err)
		return
	}
	sql = "update novel set chapter_total = chapter_total+1,chapter_current=?,words = words+? where id = ?"
	_, err = x.Exec(sql, num+1, len(content), novelId)
	if err != nil {
		logrus.Error(err)
		return
	}
}

// const (
// 	user = "root"
// 	pass = "smd123456"
// 	db   = "novel"
// 	port = "3306"
// 	// host = "60.205.191.37"
// 	host = "127.0.0.1"
// )

var x *xorm.Engine

func initDB() {
	host := os.Getenv("MYSQL_HOST")
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWD")
	db := os.Getenv("MYSQL_NOVEL_DB")
	mysql_log := os.Getenv("MYSQL_LOG") == "true"
	port := os.Getenv("MYSQL_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", user, pass, host, port, db)
	var err error
	x, err = xorm.NewEngine("mysql", dsn)
	if err != nil {
		log.Fatalf("mysql connect err:%s", err.Error())
	}
	if err := x.Ping(); err != nil {
		log.Fatalf("mysql ping err:%s", err.Error())
	}
	x.ShowSQL(mysql_log)

	x.SetMaxIdleConns(5)

	x.SetMaxOpenConns(5)
}
