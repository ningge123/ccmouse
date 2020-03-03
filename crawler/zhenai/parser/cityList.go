package parser

import (
	"github.com/PuerkitoBio/goquery"
	"io"
	"ccmouse/crawler/engine"
)

func ParseCityList(body io.ReadCloser) engine.ParseResult {
	doc, _ := goquery.NewDocumentFromReader(body)
	result := engine.ParseResult{}

	doc.Find(".city-list dd a").Each(func(i int, selection *goquery.Selection) {
		cityRequest, _ := selection.Attr("href")
		result.Requests = append(result.Requests, engine.Request{
			Url:        cityRequest,
			ParserFunc: ParseCity,
		})
	})

	return result
}