package server

import (
	"ccmouse/crawler/engine"
	"ccmouse/crawler/model"
	"ccmouse/crawler_distributed/rpcsupport"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T)  {
	const port = ":1234"

	go ServeRpc(port, "test_profile")
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(port)

	if err != nil {
		t.Fatalf("rpc client error: %v", err)
	}

	right := engine.Item{
		Url:  "http://album.zhenai.com/u/1077868794",
		Id:   "1077868794",
		Type: "zhenai",
		Payload: model.User{
			Name:       "冰之泪",
			Age:        18,
			Height:     179,
			Marriage:   "离异",
			Income:     "8001-12000元",
		},
	}

	result := ""
	err = client.Call("ItemSaverService.Save", right, &result)

	if err != nil || result != "ok" {
		t.Errorf("result is %s, err is %s", result, err)
	}
}