package parser

import (
	"crawler_v2.0/engine"
	"crawler_v2.0/model"
	"fmt"
	"regexp"
)

var pattern = regexp.MustCompile(`<script>window.__INITIAL_STATE__ = (.*?)</script>`)
var dataRe = regexp.MustCompile(`class="m-btn purple">(.+)</div>`)

var idRe = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)

// 解析器 解析用户
// name为上一级传递过来的
func ParseProfile(contents []byte, name string) engine.ParseResult {

	// 用户结构体
	profile := model.Profile{}
	profile.Name = name

	res := extractString(contents, pattern)
	fmt.Println(res)

	//for i, res := range extractString(contents, dataRe) {
	//	switch i {
	//	case 0 :
	//		profile.Marriage = string(res)
	//	case 1:
	//		ageStr := strings.Split(string(res[1]), "岁")
	//		age, err := strconv.Atoi(ageStr[0])
	//		if err == nil {
	//			profile.Age = age
	//		}
	//	case 2:
	//		profile.Constellation = string(res[1])
	//	case 3:
	//		heightStr := strings.Split(string(res[1]), "cm")
	//		height, err := strconv.Atoi(heightStr[0])
	//		if err == nil {
	//			profile.Height = height
	//		}
	//	case 4:
	//		jobAddress := strings.Split(string(res[1]), "工作地:")
	//		profile.JobAddress = jobAddress[1]
	//	case 5:
	//		salary := strings.Split(string(res[1]), "月收入:")
	//		profile.Salary = salary[1]
	//	case 6:
	//		profile.Edu = string(res[1])
	//	}
	//
	//}

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}

	return result
}

func extractString(contents []byte, re *regexp.Regexp) [][][]byte {

	match := re.FindAllSubmatch(contents, -1)

	return match
}

//type ProfileParse struct {
//	userName string
//}
//
//func (p *ProfileParse) Parse(contents []byte, url string) engine.ParseResult {
//	return ParseProfile(contents, p.userName)
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
