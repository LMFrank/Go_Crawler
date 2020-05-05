package main

import (
	"fmt"
	"github.com/jackdanger/collectlinks"
	"net/http"
	"net/url"
	"time"
)

var visited = make(map[string]bool)

func main() {
	/*
		goroutine
		拼接url，通过map去重
	*/
	url := "http://www.baidu.com"

	queue := make(chan string)
	go func() {
		queue <- url
	}()
	for uri := range queue {
		download(uri, queue)
	}
}

func download(url string, queue chan string) {
	visited[url] = true
	timeout := time.Duration(5 * time.Second)
	client := &http.Client{
		Timeout: timeout,
	}
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
		absolute := urlJoin(link, url)
		if url != " " {
			if !visited[absolute] {
				fmt.Println("parse url", absolute)
				go func() {
					queue <- absolute
				}()
			}
		}
	}
}

func urlJoin(href, base string) string {
	uri, err := url.Parse(href)
	if err != nil {
		return " "
	}
	baseUrl, err := url.Parse(base)
	if err != nil {
		return " "
	}
	return baseUrl.ResolveReference(uri).String()
}
