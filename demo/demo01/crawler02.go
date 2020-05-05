package main

import (
	"fmt"
	"github.com/jackdanger/collectlinks"
	"net/http"
)

func main() {
	url := "http://www.baidu.com/"
	download(url)
}

func download(url string) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("http get error", err)
		return
	}

	defer resp.Body.Close()

	links := collectlinks.All(resp.Body)
	for _, link := range links {
		fmt.Println("parse urls", link)
	}
}
