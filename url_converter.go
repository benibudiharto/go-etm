package etm

import (
	"net/url"
	"strings"
)

type Link struct {
	OriginalUrl string
	ExtendedParams map[string]string
	TagAttributes map[string]string
}

// Main function to generate 
func (link Link) Generate() string {
	urlObj, _ := url.Parse(link.OriginalUrl)

	combineQueryString(urlObj, link.ExtendedParams)
	combineTagAttributes(urlObj, link.TagAttributes)

	return urlObj.String()
}

// Combine additional query string to url object
func combineQueryString(urlObj *url.URL, params map[string]string) {
	q := urlObj.Query()
	for k, v := range params {
		q.Set(k, v)
	}
	urlObj.RawQuery = q.Encode()
}

// Only include tag attributes which starts with 'x-'
func combineTagAttributes(urlObj *url.URL, params map[string]string) {
	q := urlObj.Query()
	for k, v := range params {
		if strings.HasPrefix(k, "x-") {
			q.Set(k, v)
		}
	}
	urlObj.RawQuery = q.Encode()
}
