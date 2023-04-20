package main

import (
	"context"
	"fmt"

	"github.com/narumiruna/go-twse/pkg/twse"
)

func main() {
	c := twse.NewRestClient()
	data, err := c.QueryStockInfo(context.Background(), "2330")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
	fmt.Println(data.MsgArray[0].String())
}
