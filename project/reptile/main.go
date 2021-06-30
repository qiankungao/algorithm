package main

import (
	"bufio"
	"fmt"
	"github.com/anaskhan96/soup"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type SyncWriter struct {
	m      sync.Mutex
	Writer *bufio.Writer
}

func Reptile() {
	//292373
	//requestUrl := "https://m.xinwanben.com/329/"
	requestUrl := "http://www.baoziji.cc/22/22233/"
	time1 := time.Now().Unix()
	start := 1092592
	for i := start; i <= start+100; i++ {
		str := strconv.Itoa(i)
		requestUrl1 := requestUrl + str + ".html"
		GetOnePage(requestUrl1)

		requestUrl2 := requestUrl + str + "_2.html"
		GetOnePage(requestUrl2)

		requestUrl3 := requestUrl + str + "_3.html"
		GetOnePage(requestUrl3)
		fmt.Println("爬完，", requestUrl1)
	}
	time2 := time.Now().Unix()
	fmt.Println("共用时：", time2-time1)
	//GetOnePage(requestUrl)

}

func GetOnePage(requestUrl string) {
	content := SendHttpRequest(requestUrl)
	doc := soup.HTMLParse(content)
	file := OpenFile("golang.text")
	defer file.Close()

	//获取标题
	title := GetTitle(doc)
	WriteToFile(title, file)
	GetContent(doc, file)
}

func SendHttpRequest(requestUrl string) string {
	// 发送Get请求
	rsp, err := http.Get(requestUrl)
	defer rsp.Body.Close()

	if err != nil {
		log.Println(err.Error())
		return ""
	}
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		log.Println(err.Error())
		return ""
	}

	return string(body)
}

func GetTitle(doc soup.Root) string {
	subDocs := doc.FindAll("div", "class", "container")
	var title string
	for _, subDoc := range subDocs {
		h1 := subDoc.Find("h1")
		fmt.Println("h1:", h1)
		if h1.Error == nil {
			title = h1.Text()
		}
		fmt.Println("title:", title)

	}
	return title
}
func GetContent(doc soup.Root, file *os.File) {
	//获取内容
	subDocs := doc.FindAll("article", "id", "article")
	var link string
	for _, subDoc := range subDocs {
		link = subDoc.FullText()
		if link != "" {
			fmt.Println("link:", link)
		}
	}

	strs := strings.Split(link, "    ")
	for _, s := range strs {
		trimStr := strings.Trim(s, "    ")
		//fmt.Println("写入：",trimStr)
		WriteToFile(trimStr, file)
	}
}

func OpenFile(filePath string) *os.File {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		file, _ = os.Create(filePath)
	}
	return file
}

func (w *SyncWriter) Write(content string, file *os.File) {
	w.m.Lock()
	defer w.m.Unlock()
	w.Writer = bufio.NewWriter(file)
	_, _ = w.Writer.WriteString(content + "\r\n")

	//Flush将缓存的文件真正写入到文件中
	_ = w.Writer.Flush()
}

func WriteToFile(content string, file *os.File) {

	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)

	_, _ = write.WriteString(content + "\r\n")

	//Flush将缓存的文件真正写入到文件中
	_ = write.Flush()
}

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
