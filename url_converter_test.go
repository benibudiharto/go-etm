package etm

import (
	"testing"
	"net/url"
)

func TestLink_Generate(t *testing.T) {
	link := Link{
		"www.example.com?status=ok",
		map[string]string{"utm_campaign": "new-year-sale", "utm_source": "email"},
		map[string]string{"x-button": "subscribe", "x-button-color": "red", "style": "float:left;"},
	}

	result := link.Generate()
	expectedResult := "www.example.com?status=ok&utm_campaign=new-year-sale&utm_source=email&x-button=subscribe&x-button-color=red"
	if result  != expectedResult {
		t.Error("Expeceted : " + expectedResult)
		t.Error("Got       : " + result)
	}
}

func Test_CombineQueryString(t *testing.T) {
	// combine query string to url
	urlObj, _ := url.Parse("example.com?utm_campaign=new-year")
	params := map[string]string{"utm_campaign": "black-friday", "utm_source": "email"}

	combineQueryString(urlObj, params)

	if urlObj.String() != "example.com?utm_campaign=black-friday&utm_source=email" {
		t.Error("Expeceted : example.com?utm_campaign=black-friday&utm_source=email")
		t.Error("Got       : " + urlObj.String())
	}
}

func Test_CombineTagAttributes(t *testing.T) {
	urlObj, _ := url.Parse("example.com?utm_campaign=new-year")
	params := map[string]string{"x-tag-one": "1", "class": "form-control", "x-tag-two": "two", "width": "500px"}

	combineTagAttributes(urlObj, params)

	expectedResult := "example.com?utm_campaign=new-year&x-tag-one=1&x-tag-two=two"
	if urlObj.String() != expectedResult {
		t.Error("Expeceted : " + expectedResult)
		t.Error("Got       : " + urlObj.String())
	}
}
