package main

import (
	"flag"
	"fmt"
	"go-crawler/crawler"
	"go-crawler/files"
	"log"
	"time"
)

const pageCountDefault = 3 // 默认总页数为3

const urlFormat = "https://xskydata.jobs.feishu.cn/school/?current=%d"

const selector = "#bd > section > section > main > div > div > div.content__bb7170 > div.rightBlock.rightBlock__bb7170 > div.borderContainer__bb7170 > div.listItems__bb7170"

func main() {
	// 解析命令作为页数参数
	pageCount := flag.Int("page_count", pageCountDefault, "The count of page to crawl")
	flag.Parse()
	fmt.Println("爬取页数为: ", *pageCount)
	res, err := crawler.TotalCrawler(urlFormat, selector, *pageCount)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = files.WriteIntoJSON(res)
	if err != nil {
		log.Fatal("Fail to write the data into the json file")
		return
	}

	fmt.Println("爬取成功，请打开data.json文件查看")
	time.Sleep(2 * time.Second) // 阻塞页面查看输出结果
}
