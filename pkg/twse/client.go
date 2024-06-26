package twse

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	defaultTimeout = 15 * time.Second
	baseURL        = "https://mis.twse.com.tw"
)

type Client struct {
	client  *http.Client
	baseURL *url.URL
}

func NewClient() *Client {
	u, err := url.Parse(baseURL)
	if err != nil {
		panic(err)
	}

	return &Client{
		client: &http.Client{
			Timeout: defaultTimeout,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		},
		baseURL: u,
	}
}

func (c *Client) NewRequest(ctx context.Context, method string, refURL string, params url.Values) (*http.Request, error) {
	rel, err := url.Parse(refURL)
	if err != nil {
		return nil, err
	}

	if params != nil {
		rel.RawQuery = params.Encode()
	}

	u := c.baseURL.ResolveReference(rel)

	req, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	return req, nil
}

func (c *Client) QueryStockInfo(ctx context.Context, symbols ...string) (*Response, error) {
	params := url.Values{}
	params.Add("ex_ch", buildExCh(symbols...))
	params.Add("json", "1")
	params.Add("delay", "0")
	params.Add("_", strconv.FormatInt(time.Now().UnixMilli(), 10))

	req, err := c.NewRequest(ctx, "GET", "/stock/api/getStockInfo.jsp", params)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data Response
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

func buildExCh(symbols ...string) string {
	var slice []string
	for _, s := range symbols {
		slice = append(slice, fmt.Sprintf("tse_%s.tw", s))
		slice = append(slice, fmt.Sprintf("otc_%s.tw", s))
	}
	return strings.Join(slice, "|")
}
