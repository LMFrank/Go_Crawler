package model

import "strconv"

type Bookdetails struct {
	Bookname string
	Author   string
	Press    string
	Pages    int
	Price    string
	Score    string
	Intro    string
}

func (b Bookdetails) String() string {
	return "\n书名：" + b.Bookname + "\n作者：" + b.Author + "\n出版社：" + b.Press + "\n页数：" +
		strconv.Itoa(b.Pages) + "\n价格：" + b.Price + "\n评分：" + b.Score + "\n简介：" + b.Intro
}
