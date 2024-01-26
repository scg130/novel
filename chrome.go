package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

// 定义一个爬虫结构体
type Crawler struct {
}

// 面向对象思想，将配置定义成方法
func (c Crawler) config() (opts []selenium.ServiceOption, caps selenium.Capabilities) {
	opts = []selenium.ServiceOption{}
	caps = selenium.Capabilities{
		"browserName": "chrome",
	}
	// 禁止加载图片，加快渲染速度
	imagCaps := map[string]interface{}{
		//"profile.managed_default_content_settings.images": 2,
	}
	chromeCaps := chrome.Capabilities{
		Prefs: imagCaps,
		Path:  "",
		Args: []string{
			//"--headless", // 设置Chrome无头模式
			"--no-sandbox",
			"--user-agent=Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/604.4.7 (KHTML, like Gecko) Version/11.0.2 Safari/604.4.7", // 模拟user-agent，防反爬
		},
	}
	caps.AddChrome(chromeCaps)
	return opts, caps
}

// 爬虫启动
func (c Crawler) startChrome(path string, opts []selenium.ServiceOption, caps selenium.Capabilities) (*selenium.Service, selenium.WebDriver) {
	// 启动chromedriver，端口号可自定义
	service, err := selenium.NewChromeDriverService(path, 9999, opts...)
	if err != nil {
		log.Printf("Error starting the ChromeDriver server: %v", err)
	}

	// 调起chrome浏览器
	webDriver, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", 9999))
	if err != nil {
		panic(err)
	}
	//defer service.Stop()   // 停止chromedriver
	//defer webDriver.Quit() // 关闭浏览器
	return service, webDriver
}

func main() {
	driverPath := "chromedriver" //准备工作中下载driver
	crawler := Crawler{}
	opts, caps := crawler.config()
	service, driver := crawler.startChrome(driverPath, opts, caps)

	// targetUrl := "https://www.baidu.com"
	// err := driver.Get(targetUrl)
	// if err != nil {
	// 	panic(fmt.Sprintf("Failed to load page: %s\n", err))
	// }

	// inputElement, _ := driver.FindElement("id", "kw")
	// inputElement.SendKeys("万古神帝最新章节")
	// inputElement.Submit()

	for {
		windows, err := driver.WindowHandles()
		if err != nil {
			panic(err)
		}
		fmt.Println(windows)
		driver.SwitchWindow(windows[0])

		result, err := driver.PageSource()
		// 将结果写入goquery中，以便用css选择器过滤标签
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(result))
		if err != nil {
			panic(fmt.Sprintf("Failed: %s\n", err))
		}

		doc.Find(".book_last").Find("dd").Each(func(i int, s *goquery.Selection) {
			text := s.Find("a").Text()
			href, _ := s.Find("a").Attr("href")
			fmt.Println(text, href)
		})

		time.Sleep(time.Duration(5) * time.Second)
	}

	service.Stop() // 停止chromedriver
	driver.Quit()  // 关闭浏览器

}
