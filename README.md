# go-crawler
A crawler program with Go for XSKY

### 项目描述
基于Go实现的爬虫，爬取招聘网站的所有岗位，并把结果转为`json`写入到根目录的`data.json`文件下

### 项目运行
命令格式：\
1.exe运行
```shell
./main.exe -page_count [number] 
```
其中`number`为数字，表示要爬取的页数，默认为3页，其中每页最多有10个职位信息。
也可以不指定`page_count`参数，默认为3页:
```shell
./main.exe
```
2.go运行
```shell
go run main.go -page_count [number]
```