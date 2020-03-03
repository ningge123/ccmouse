package parser

import (
	"github.com/PuerkitoBio/goquery"
	"io"
	"ccmouse/crawler/engine"
	"strings"
)

func ParseCity(body io.ReadCloser) engine.ParseResult {
	var result engine.ParseResult
	doc, _ := goquery.NewDocumentFromReader(body)

	doc.Find(".g-list .list-item .content").Each(func(i int, selection *goquery.Selection) {
		var name string
		var url string
		var gender string
		selection.Find("tr").First().Find("a").Each(func(i int, selection *goquery.Selection) {
			name = selection.Text()
			url, _ = selection.Attr("href")
		})
		selection.Find("tr").First().Next().Children().First().Each(func(i int, selection *goquery.Selection) {
			gender = selection.Text()
			gender = strings.ReplaceAll(gender, "性别：", "")
		})

		result.Requests = append(result.Requests, engine.Request{
			Url:        url,
			ParserFunc: func(closer io.ReadCloser) engine.ParseResult {
				return ParseUser(url, name, gender, closer)
			},
		})
	})

	doc.Find(".f-pager .m-page a").Each(func(i int, selection *goquery.Selection) {
		url, _ := selection.Attr("href")
		bool := strings.Contains(url, "http://www.zhenai.com/zhenghun")

		if bool {
			result.Requests = append(result.Requests, engine.Request{
				Url:        url,
				ParserFunc: ParseCity,
			})
		}
	})

	return result
}
