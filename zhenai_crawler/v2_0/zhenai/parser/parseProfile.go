package parser

import (
	"cralwer_v2.0/engine"
	"cralwer_v2.0/model"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([\d]+)岁</div>`)
var marry = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>(已婚)</div>`)
var constellation = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>(天秤座)</div>`)
var height = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([\d]+)cm</div>`)
var weight = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([\d]+)kg</div>`)
var salary = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>月收入:([^<]+)</div>`)

var idRe = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)

// 解析器 解析用户
// name为上一级传递过来的
func ParseProfile(contents []byte, name string) engine.ParseResult {

	// ioutil.WriteFile("test.html",contents,0x777)
	// 用户结构体
	profile := model.Profile{}
	profile.Name = name

	// 年龄   string转换为int
	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err == nil {
		profile.Age = age
	}
	// 身高
	height, err := strconv.Atoi(extractString(contents, height))
	if err == nil {
		profile.Height = height
	}
	// 体重
	weight, err := strconv.Atoi(extractString(contents, weight))
	if err == nil {
		profile.Weight = weight
	}

	// 薪水
	profile.Salary = extractString(contents, salary)

	// 星座
	profile.Constellation = extractString(contents, constellation)
	if extractString(contents, marry) == "" {
		profile.Marry = "未婚"
	} else {
		profile.Marry = "已婚"
	}

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}

	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {

	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}

//type ProfileParse struct {
//	userName string
//}
//
//func (p *ProfileParse) Parse(contents []byte, url string) engine.ParseResult {
//	return ParseProfile(contents, url, p.userName)
//}
//
//func (p *ProfileParse) Serialize() (name string, args interface{}) {
//	return "ProfileParse", p.userName
//}
//
//func NewprofileParse(name string) *ProfileParse {
//	return &ProfileParse{
//		userName: name,
//	}
//}
