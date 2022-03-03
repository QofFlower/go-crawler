package files

// @Author: Hananoq

import (
	"encoding/json"
	"fmt"
	"go-crawler/crawler"
	"io/ioutil"
)

func WriteIntoJSON(jobs []crawler.Job) error {
	// 解析结构体
	output, err := json.MarshalIndent(&struct {
		Count int           `json:"count"`
		Jobs  []crawler.Job `json:"jobs"`
	}{len(jobs), jobs}, "", "\t\t")
	if err != nil {
		fmt.Println("转换失败")
		return err
	}
	// 写入文件
	err = ioutil.WriteFile("data.json", output, 0644)
	if err != nil {
		fmt.Println("写文件失败", err)
		return err
	}
	return nil
}
