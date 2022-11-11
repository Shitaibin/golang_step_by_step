package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"time"
)

type Simple struct {
	endpoint string // http://ip:port/v1
	*http.Client
}

func NewClient(endpoint string, transport http.RoundTripper) (*Simple, error) {
	if _, err := url.Parse(endpoint); err != nil {
		return nil, err
	}

	return &Simple{
		endpoint: endpoint,
		Client:   &http.Client{Transport: transport},
	}, nil
}

const (
	HTTP_GET    = "GET"
	HTTP_POST   = "POST"
	HTTP_DELETE = "DELETE"
)

type Request struct {
	LeftUrl    string
	HttpMethod string
	Params     map[string]string
	Body       interface{}
}

func (s *Simple) Quest(ctx context.Context, req *Request, response interface{}) error {
	// build request
	urlPath := path.Join(s.endpoint, req.LeftUrl)
	u, err := url.Parse(urlPath)
	if err != nil {
		return fmt.Errorf("parse url error: %v", err)
	}

	q := u.Query()
	q.Set("timestamp", strconv.FormatInt(time.Now().Unix(), 10))
	q.Set("sender", "dabin")
	for k, v := range req.Params {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()

	bodyBuf, err := json.Marshal(req.Body)
	if err != nil {
		return fmt.Errorf("marshal body error: %v", err)
	}

	body := bytes.NewReader(bodyBuf)
	httpReq, err := http.NewRequestWithContext(ctx, req.HttpMethod, u.String(), body)
	if err != nil {
		return fmt.Errorf("create request error: %v", err)
	}

	// send request
	resp, err := s.Do(httpReq)
	if err != nil {
		return fmt.Errorf("send request error: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("http status not ok: %v", resp.StatusCode)
	}

	// unmarshal response
	if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
		return fmt.Errorf("decode response body error: %v", err)
	}
	defer func() {
		resp.Body.Close()
	}()

	return nil
}
