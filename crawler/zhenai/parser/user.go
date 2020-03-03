package parser

import (
	"github.com/PuerkitoBio/goquery"
	"io"
	"regexp"
	"ccmouse/crawler/engine"
	"ccmouse/crawler/model"
	"strconv"
	"strings"
)

func ParseUser(url string, name string, gender string, body io.ReadCloser) engine.ParseResult {
	var user model.User
	user.Name = name
	user.Gender = gender
	doc, _ := goquery.NewDocumentFromReader(body)

	doc.Find(".des.f-cl").Each(func(i int, selection *goquery.Selection) {
		desc := selection.Text()

		desc = strings.ReplaceAll(desc, " ", "")
		descSlice := strings.Split(desc, "|")

		user.City = descSlice[0]
		user.Education = descSlice[2]
		user.Marriage = descSlice[3]
		// 提取出年龄身高
		reg, _ := regexp.Compile(`([\d]+)`)
		age := reg.FindString(descSlice[1])
		height := reg.FindString(descSlice[4])
		user.Age, _ = strconv.Atoi(age)
		user.Height, _ = strconv.Atoi(height)
		user.Income = descSlice[5]
	})

	var id string
	doc.Find(".info .id").Each(func(i int, selection *goquery.Selection) {
		text := selection.Text()

		id = strings.ReplaceAll(text, "ID：", "")
	})

	var item []engine.Item
	return engine.ParseResult{
		Items: append(item, engine.Item{
			Url:  url,
			Id:   id,
			Type: "zhenai",
			Payload: user,
		}),
	}
}
