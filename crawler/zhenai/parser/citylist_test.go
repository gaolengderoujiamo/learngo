package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("cityparser_test_data.html")

	if err != nil {
		panic(err)
	}

	parserResult := ParseCityList(contents)

	const resultSize = 470

	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}

	expectedCities := []string{
		"阿坝", "阿克苏", "阿拉善盟",
	}

	if len(parserResult.Requests) != resultSize {
		t.Errorf("Result should have %d requests; but had %d",
			resultSize, len(parserResult.Requests))
	}

	for i, url := range expectedUrls {
		if parserResult.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s; but was %s",
				i, url, parserResult.Requests[i].Url)
		}
	}

	if len(parserResult.Items) != resultSize {
		t.Errorf("Result should have %d items; but had %d",
			resultSize, len(parserResult.Items))
	}

	for i, city := range expectedCities {
		if parserResult.Items[i].(string) != city {
			t.Errorf("expected city #%d: %s; but was %s",
				i, city, parserResult.Items[i].(string))
		}
	}
}
