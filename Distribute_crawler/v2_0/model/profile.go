package model

import (
	"encoding/json"
	"strconv"
)

type Profile struct {
	Bookname string
	Author   string
	Press    string
	Pages    int
	Price    string
	Score    float64
	ISBN     int
}

func (p Profile) String() string {
	return "\n书名：" + p.Bookname + "\n作者：" + p.Author + "\n出版社：" + p.Press + "\n页数：" + strconv.Itoa(p.Pages) +
		"\n价格：" + p.Price + "\n评分：" + strconv.FormatFloat(p.Score, 'f', 1, 64) + "\nISBN：" + strconv.Itoa(p.ISBN)
}

func FromJsonObj(o interface{}) (Profile, error) {
	var profile Profile
	s, err := json.Marshal(o)
	if err != nil {
		return profile, err
	}

	err = json.Unmarshal(s, &profile)
	return profile, err
}
