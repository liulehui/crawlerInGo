package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/liulehui/crawler/concurrant/config"
	rpcnames "github.com/liulehui/crawler/distributed/config"
	"github.com/liulehui/crawler/distributed/rpcsupport"
	"github.com/liulehui/crawler/distributed/worker"
)

func TestCrawlService(t *testing.T) {
	const host = ":9000"
	go rpcsupport.ServeRpc(
		host, worker.CrawlService{})
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	// TODO: Use a fake fetcher to handle the url.
	// So we don't get data from zhenai.com
	req := worker.Request{
		Url: "http://album.zhenai.com/u/108906739",
		Parser: worker.SerializedParser{
			Name: config.ParseProfile,
			Args: "安静的雪",
		},
	}
	var result worker.ParseResult
	err = client.Call(
		rpcnames.CrawlServiceRpc, req, &result)

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(result)
	}

	// TODO: Verify results
}
