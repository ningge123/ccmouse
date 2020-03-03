package engine

import "io"

type Request struct {
	Url        string
	ParserFunc func(io.ReadCloser) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Url     string
	Id      string
	Type    string
	Payload interface{}
}