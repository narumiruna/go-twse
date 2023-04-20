package main

import (
	"context"
	"fmt"

	"github.com/narumiruna/go-twse/pkg/twse"
)

func main() {
	c := twse.NewRestClient()
	data, err := c.QueryStockInfo(context.Background(), "2330 0050")
	if err != nil {
		panic(err)
	}

	for _, info := range data.MsgArray {
		fmt.Println(info.String())
	}
}
