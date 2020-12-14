package config

const (
	// Parser names
	ParseBookTag    = "ParseBookTag"
	ParseBookList   = "ParseBookList"
	ParseBookDetail = "ParseBookDetail"
	NilParser       = "NilParser"

	// Service port
	ItemSaverPort = 1234
	WorkerPort0   = 9000

	//ElasticSearch
	ElasticIndex = "douban_book"

	// RPC Endpoints
	ItemSaverRpc    = "ItemSaverService.Save"
	CrawlServiceRpc = "CrawlService.Process"
)
