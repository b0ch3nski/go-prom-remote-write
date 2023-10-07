package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/b0ch3nski/go-prom-remote-write/model"

	"github.com/golang/snappy"
	"google.golang.org/protobuf/proto"
)

type Client interface {
	Write(context.Context, []*model.TimeSeries) error
}

var _ Client = (*client)(nil)

type client struct {
	endpoint string
	username string
	password string
	timeout  time.Duration
	hcl      *http.Client
}

// NewClient creates new Prometheus remote write client.
func NewClient(endpoint string) *client {
	defaultTimeout := 5 * time.Second

	return &client{
		endpoint: endpoint,
		timeout:  defaultTimeout,
		hcl:      &http.Client{Timeout: defaultTimeout},
	}
}

func (c *client) WithTimeout(timeout time.Duration) *client {
	c.timeout = timeout
	return c
}

func (c *client) WithAuthBasic(username, password string) *client {
	c.username, c.password = username, password
	return c
}

func (c *client) WithHttpClient(cl *http.Client) *client {
	c.hcl = cl
	return c
}

// Write sends time series data to Prometheus using remote write mechanism.
// See specification: https://prometheus.io/docs/concepts/remote_write_spec
func (c *client) Write(ctx context.Context, series []*model.TimeSeries) error {
	seriesProto, errMarshal := proto.Marshal(&model.WriteRequest{Timeseries: series})
	if errMarshal != nil {
		return fmt.Errorf("failed marshaling remote write request: %w", errMarshal)
	}
	seriesSnappy := snappy.Encode(nil, seriesProto)

	ctxReq, cancelReq := context.WithTimeout(ctx, c.timeout)
	defer cancelReq()

	req, errReq := http.NewRequestWithContext(ctxReq, http.MethodPost, c.endpoint, bytes.NewBuffer(seriesSnappy))
	if errReq != nil {
		return fmt.Errorf("failed creating http request: %w", errReq)
	}
	if c.username != "" && c.password != "" {
		req.SetBasicAuth(c.username, c.password)
	}
	req.Header.Set("Content-Type", "application/x-protobuf")
	req.Header.Set("Content-Encoding", "snappy")
	req.Header.Set("X-Prometheus-Remote-Write-Version", "0.1.0")

	resp, errResp := c.hcl.Do(req)
	if errResp != nil {
		return fmt.Errorf("failed executing http request: %w", errResp)
	}
	defer resp.Body.Close()

	if code := resp.StatusCode; code/100 != 2 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("unsuccessful http status code: %d: %s", code, string(body))
	}
	return nil
}
