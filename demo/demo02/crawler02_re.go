package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func main() {
	/*
		正则匹配
		<a href="/tag/小说" class="tag">小说</a>
		将 /tag/小说 取出，用于url拼接
	*/

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://book.douban.com/", nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error status code:%d\n", resp.StatusCode)
	}

	//对网站编码做检测
	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())

	result, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}

	//fmt.Println(string(result))
	parseContent(result)
}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("fetch error: %v", err)
		return unicode.UTF8
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

func parseContent(content []byte) {
	//<a href="/tag/小说" class="tag">小说</a>
	re := regexp.MustCompile(`<a href="([^"]+)" class="tag">([^"]+)</a>`)
	match := re.FindAllSubmatch(content, -1)

	for _, m := range match {
		fmt.Printf("m[0]: %s m[1]: %s m[2]: %s\n", m[0], m[1], m[2])
		fmt.Printf("url:%s\n", "https://book.douban.com"+string(m[1]))
	}
}
