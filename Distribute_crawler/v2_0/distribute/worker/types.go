package worker

import (
	"crawler_v2.0/distribute/config"
	"crawler_v2.0/doubanbook/parser"
	"crawler_v2.0/engine"
	"errors"
	"fmt"
	"log"
)

type SerializedParser struct {
	Name string
	Args interface{}
}

type Request struct {
	Url    string
	Parser SerializedParser
}

type ParseResult struct {
	Items    []engine.Item
	Requests []Request
}

func SerializeRequest(r engine.Request) Request {
	name, args := r.Parser.Serialize()
	return Request{
		Url: r.Url,
		Parser: SerializedParser{
			Name: name,
			Args: args,
		},
	}
}

func SerializeResult(r engine.ParseResult) ParseResult {
	result := ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		result.Requests = append(result.Requests, SerializeRequest(req))
	}
	return result
}

func DeserializeRequest(r Request) (engine.Request, error) {
	parser, err := deserializeParser(r.Parser)
	if err != nil {
		return engine.Request{}, err
	}
	return engine.Request{
		Url:    r.Url,
		Parser: parser,
	}, nil
}

func DeserializeResult(r ParseResult) engine.ParseResult {
	result := engine.ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		engineReq, err := DeserializeRequest(req)
		if err != nil {
			log.Printf("error deserializing "+"request: %v", err)
			continue
		}
		result.Requests = append(result.Requests, engineReq)
	}
	return result
}

func deserializeParser(p SerializedParser) (engine.Parser, error) {
	switch p.Name {
	case config.ParseBookTag:
		return engine.NewFuncParser(parser.ParseBookTag, config.ParseBookTag), nil
	case config.ParseBookList:
		return engine.NewFuncParser(parser.ParseBookList, config.ParseBookList), nil
	case config.NilParser:
		return engine.NilParser{}, nil
	case config.ParseBookDetail:
		if bookName, ok := p.Args.(string); ok {
			return parser.NewBookDetailParser(bookName), nil
		} else {
			return nil, fmt.Errorf("invalid "+"arg: %v", p.Args)
		}
	default:
		return nil, errors.New("unknown parser name")
	}
}
