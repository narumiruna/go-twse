package twse

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const defaultTimeout = 15 * time.Second
const baseApiUrl = "https://mis.twse.com.tw"

type RestClient struct {
	client  *http.Client
	baseURL *url.URL
}

func NewRestClient() *RestClient {
	u, err := url.Parse(baseApiUrl)
	if err != nil {
		panic(err)
	}

	return &RestClient{
		client: &http.Client{
			Timeout: defaultTimeout,
		},
		baseURL: u,
	}
}

func (c *RestClient) NewRequest(ctx context.Context, method string, refURL string, params url.Values) (*http.Request, error) {
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

func (c *RestClient) QueryStockInfo(ctx context.Context, symbol string) (*Response, error) {
	params := url.Values{}
	params.Add("ex_ch", fmt.Sprintf("tse_%s.tw", symbol))
	params.Add("json", "1")
	params.Add("delay", "0")
	params.Add("_", strconv.FormatInt(time.Now().UnixMilli(), 10))

	req, err := c.NewRequest(ctx, "GET", "/stock/api/getStockInfo.jsp", params)
	if err != nil {
		return nil, err
	}
	fmt.Println(req.URL)

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
