package main

import (
	"context"
	"fmt"

	"github.com/narumiruna/go-twse/pkg/twse"
)

func main() {
	symbols := []string{"2330", "0050", "006208", "2254", "6526", "6901", "6757", "4569", "6902", "6657", "3536", "8480",
		"3383",
		"2254",
		"6526",
		"6901",
		"6757",
		"4569",
		"6902",
		"6657",
		"2432",
		"6861",
		"6606",
	}
	c := twse.NewClient()
	data, err := c.QueryStockInfo(context.Background(), symbols...)
	if err != nil {
		panic(err)
	}

	for _, info := range data.MsgArray {
		fmt.Println(info)
	}
}
