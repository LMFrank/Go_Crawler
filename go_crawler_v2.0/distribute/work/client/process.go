package client

import (
	"go_crawler/distribute/rpcsupport"
	"go_crawler/distribute/work"
	"go_crawler/engine"
)

func CreateProcess() (engine.Processor, error) {
	client, err := rpcsupport.NewClient(":1235")

	if err != nil {
		return nil, err
	}

	return func(req engine.Request) (engine.ParseResult, error) {
		sReq := work.SerializeRequest(req)

		var sResult work.ParseResult

		err := client.Call("CrawlService.Process", sReq, &sResult)

		if err != nil {
			return engine.ParseResult{}, nil
		}
		return work.DeserializeResult(sResult), nil

	}, nil

}
