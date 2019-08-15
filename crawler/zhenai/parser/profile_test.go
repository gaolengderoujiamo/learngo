package parser

import (
	"io/ioutil"
	"learngo/crawler/model"
	"testing"
)

func TestParserProfile(t *testing.T) {
	contents, e := ioutil.ReadFile("profile_test_data.html")
	if e != nil {
		panic(e)
	}

	extractString(contents, profileRe, 5)

	result := ParserProfile(contents, "劣酒灼心")

	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 element; but was %v", result.Items)
	}

	profile := result.Items[0].(model.Profile)

	expected := model.Profile{
		Name:      "劣酒灼心",
		Gender:    "男",
		Age:       25,
		Height:    165,
		Income:    "5001-8000元",
		Marriage:  "未婚",
		Education: "中专",
		Area:      "阿坝",
	}

	if profile != expected {
		t.Errorf("expected %v; but was %v", expected, profile)
	}
}
