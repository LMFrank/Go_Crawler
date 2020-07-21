package model

import "strconv"

type Profile struct {
	Name          string
	Marriage      string
	Age           int
	Constellation string
	Height        int
	JobAddress    string
	Salary        string
	Edu           string
}

func (p Profile) String() string {
	return p.Name + " " + p.Marriage + " " + strconv.Itoa(p.Age) + "岁 " +
		strconv.Itoa(p.Height) + "cm " + p.JobAddress + " 月收入：" + p.Salary + " 学历：" + p.Edu
}
