package etm

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
	"html"
)

type Etm struct {
	Html string
	QueryParams map[string]string
}

func (etm Etm) String() string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(etm.Html))

	if err != nil {
		log.Fatal(err)
	}

	// Iterate hyperlinks
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists == false {
			return;
		}

		// get all
		node := s.Nodes[0]
		attrMap := map[string]string{}
		for _, attr := range node.Attr {
			attrMap[attr.Key] = attr.Val
		}

		link := Link{href, etm.QueryParams, attrMap}

		s.SetAttr("href", html.UnescapeString(link.Generate()))
	})

	resultHtml, err := doc.Html()

	if err != nil {
		log.Fatal(err)
	}
	return resultHtml
}

func Convert(html string, params map[string]string) string {
	etm := Etm{html, params}
	return etm.String()
}
