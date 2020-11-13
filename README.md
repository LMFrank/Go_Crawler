## Go_Crawler

Go写的一些爬虫项目

Python爬虫项目地址：[ https://github.com/LMFrank/CrawlerProject ]( https://github.com/LMFrank/CrawlerProject )

### demo01

初始Go爬虫，利用原生的`net/http`库爬取

### demo02

1. 对网页`charset`的检测，并将其统一编码为utf-8
2. 正则表达式匹配，用于url拼接

### Go_crawler_v1.0

以爬取[豆瓣读书]( https://book.douban.com/ )为例，构建爬虫项目，~~后期尝试改造为分布式爬虫~~。

已改造为并发版，数据存储使用`elasticsearch`

### Go_crawler_v2.0

在v1.0的版本上改造了并发结构，尝试拆分为微服务，使用内置的rpc进行通信

### Distribute_cralwer

Go_crawler_v2.0中的豆瓣网爬虫项目

在本项目中重构了代码，加入了存储模块（ES），完善了框架

具体请查看：[README](https://github.com/LMFrank/Go_Crawler/tree/master/Distribute_crawler)

