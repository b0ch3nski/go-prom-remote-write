# go-prom-remote-write

Bare minimum Prometheus Remote-Write client, based on Proto files acquired from upstream [repository][repo].
All dependencies were cut down with only `snappy` and `protobuf` remaining.

For a code generation procedure, see included [Makefile](Makefile).

Simple client fulfills the Remote-Write [specification][spec] with couple small additions, like ability to handle
timeouts and Basic authentication.

[repo]: https://github.com/prometheus/prometheus/tree/main/prompb
[spec]: https://prometheus.io/docs/concepts/remote_write_spec

## install

```
go get github.com/b0ch3nski/go-prom-remote-write
```

## example

```go
promClient := client.
	NewClient("http://localhost:9090/api/v1/write").
	WithAuthBasic("username", "password").
	WithTimeout(3 * time.Second).
	WithHttpClient(&http.Client{Transport: &http.Transport{MaxConnsPerHost: 0}})

series := []*model.TimeSeries{
	{
		Samples: []*model.Sample{
			{
				Value:     321.123,
				Timestamp: time.Now().UTC().UnixMilli(),
			},
		},
		Labels: []*model.Label{
			{
				Name:  "__name__",
				Value: "test_metric",
			},
		},
	},
}

if err := promClient.Write(context.Background(), series); err != nil {
	panic(err)
}
```
