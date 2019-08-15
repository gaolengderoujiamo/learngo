package parser

import (
	"learngo/crawler/engine"
	"learngo/crawler/model"
	"regexp"
	"strconv"
	"strings"
)

// 阿坝 | 25岁 | 中专 | 未婚 | 165cm | 5001-8000元
var profileRe = regexp.MustCompile(`<div class="des f-cl" data-v-3c42fade>([^\|]+)\|([^\|]+)\|([^\|]+)\|([^\|]+)\|([^\|]+)\|([^<]+)</div>`)
var genderRe = regexp.MustCompile(`<a href="http://www.zhenai.com/zhenghun/[\w]+/([\w]+)">[^<]*士征婚</a>`)

func ParserProfile(contents []byte, name string) engine.ParserResult {
	profile := model.Profile{}

	profile.Name = name
	genderStr := extractString(contents, genderRe, 1)
	if genderStr == "nv" {
		profile.Gender = "女"
	} else if genderStr == "nan" {
		profile.Gender = "男"
	} else {
		profile.Gender = "未知"
	}
	ageRune := []rune(extractString(contents, profileRe, 2))
	if len(ageRune) > 1 {
		if age, err := strconv.Atoi(string(ageRune[:len(ageRune)-1])); err == nil {
			profile.Age = age
		}
	}
	heightRune := []rune(extractString(contents, profileRe, 5))
	if len(heightRune) > 2 {
		if height, err := strconv.Atoi(string(heightRune[:len(heightRune)-2])); err == nil {
			profile.Height = height
		}
	}
	profile.Income = string(extractString(contents, profileRe, 6))
	profile.Marriage = string(extractString(contents, profileRe, 4))
	profile.Education = string(extractString(contents, profileRe, 3))
	profile.Area = string(extractString(contents, profileRe, 1))

	result := engine.ParserResult{
		Items: []interface{}{profile},
	}

	return result
}

func extractString(contents []byte, re *regexp.Regexp, index int) string {
	match := re.FindSubmatch(contents)

	if len(match) >= 1+index {
		return strings.Trim(string(match[index]), " ")
	} else {
		return ""
	}
}
