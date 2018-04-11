package parser

import (
	"crawler/model"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")

	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents, "")

	if len(result.Items) != 1 {
		t.Errorf("Result should contain 1 element; but was %v", result.Items)
	}

	profile := result.Items[0].(model.Profile)
	expected := model.Profile{
		Age:        1,
		Height:     1,
		Weight:     2,
		Income:     "123",
		Gender:     "女",
		Name:       "米琪",
		Xinzuo:     "aaa",
		Occupation: "bbb",
		Marriage:   "ccc",
		House:      "ddd",
		Hokou:      "eee",
		Education:  "fff",
		Car:        "fff",
	}

	if profile != expected {
		t.Errorf("expected %v; but was %v", expected, profile)
	}
}
