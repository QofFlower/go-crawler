package crawler

// @Author: Hananoq
import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"log"
	"strings"
	"time"
)

const ( // 岗位类型标识码
	JobTypeName = iota
	JobTypeBaseAndType
	JobTypeDescription
)

type SelectorData struct {
	Type     int
	Selector string
}

type Job struct {
	Name        string `json:"name"`        // 岗位名称
	Base        string `json:"base"`        // 岗位地点
	Type        string `json:"type"`        // 岗位类型
	Description string `json:"description"` // 岗位描述
}

var sData = []*SelectorData{
	{JobTypeName, ".positionItem-title-text"},
	{JobTypeBaseAndType, ".subTitle__bb7170,.positionItem-subTitle"},
	{JobTypeDescription, ".jobDesc__bb7170,.positionItem-jobDesc"}}

// GetHttpHtmlContent 获取网站上爬取的数据
func GetHttpHtmlContent(url string, selector string, sel interface{}) (string, error) {
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", true), // debug使用
		chromedp.Flag("blink-settings", "imagesEnabled=false"),
		chromedp.UserAgent(`Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`),
	}
	//初始化参数，先传一个空的数据
	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)

	c, _ := chromedp.NewExecAllocator(context.Background(), options...)

	// create context
	chromeCtx, cancel := chromedp.NewContext(c, chromedp.WithLogf(log.Printf))
	// 执行一个空task, 用提前创建Chrome实例
	_ = chromedp.Run(chromeCtx, make([]chromedp.Action, 0, 1)...)

	//创建一个上下文，具有超时时间，用于等待页面加载完成，爬取js生产的数据
	timeoutCtx, cancel := context.WithTimeout(chromeCtx, 30*time.Second)
	defer cancel()

	var htmlContent string
	err := chromedp.Run(timeoutCtx,
		chromedp.Navigate(url),
		chromedp.WaitVisible(selector),
		chromedp.OuterHTML(sel, &htmlContent, chromedp.ByJSPath),
	)
	if err != nil {
		//log.Fatal("Run err : %v\n", err)
		return "", err
	}

	return htmlContent, nil
}

// JoinSelector selector过滤规则拼接
func JoinSelector(selectorData ...*SelectorData) string {
	var strs []string
	for _, v := range selectorData {
		strs = append(strs, v.Selector)
	}
	temp := strings.Join(strs, ",")
	return temp[:len(temp)-1]
}

// GetSpecialData 过滤出具体的数据
func GetSpecialData(htmlContent string, selectorData ...*SelectorData) ([]Job, error) {
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var tempJob Job
	var res []Job
	dom.Find(JoinSelector(selectorData...)).Each(func(i int, selection *goquery.Selection) {
		if i == 0 || i%3 == JobTypeName {
			tempJob.Name = selection.Text()
		} else if i%3 == JobTypeBaseAndType {
			runes := []rune(selection.Text()) // 中文字节数不一样，直接截取会乱码
			tempJob.Base, tempJob.Type = string(runes[:2]), string(runes[2:])
		} else if i%3 == JobTypeDescription {
			tempJob.Description = selection.Text()
		}
		if (i+1)%3 == 0 {
			res = append(res, tempJob)
		}
	})
	return res, nil
}

func TotalCrawler(urlFormat, selector string, pageCount int) ([]Job, error) {
	var res []Job
	for i := 0; i < pageCount; i++ {
		currentUrl := fmt.Sprintf(urlFormat, i+1)
		if jobs, err := crawler(currentUrl, selector); err != nil {
			log.Fatal("Fail to filter the data: ", err)
			return nil, err
		} else {
			res = append(res, jobs...)
		}
	}
	return res, nil
}

// 单页爬取
func crawler(url, selector string) ([]Job, error) {
	param := `document.querySelector("body")`
	html, _ := GetHttpHtmlContent(url, selector, param)

	return GetSpecialData(html, sData...)
}
