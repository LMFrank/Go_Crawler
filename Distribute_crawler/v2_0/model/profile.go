package model

import "strconv"

type Profile struct {
	Bookname string
	Author   string
	Press    string
	Pages    int
	Price    string
	Score    float64
	Intro    string
}

func (b Profile) String() string {
	return "\n书名：" + b.Bookname + "\n作者：" + b.Author + "\n出版社：" + b.Press + "\n页数：" +
		strconv.Itoa(b.Pages) + "\n价格：" + b.Price + "\n评分：" + strconv.FormatFloat(b.Score, 'f', 1, 64) + "\n简介：" + b.Intro
}
