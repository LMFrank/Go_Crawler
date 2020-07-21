package test

import (
	"cralwer_v2.0/zhenai/parser"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}
	result := parser.ParseCityList(contents)
	fmt.Println(result.Items)

	const resultSize = 470
	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d "+"requests, but has %d", resultSize, len(result.Requests))
	}

	if len(result.Items) != resultSize {
		t.Errorf("result should have %d "+"requests, but has %d", resultSize, len(result.Items))
	}
}
