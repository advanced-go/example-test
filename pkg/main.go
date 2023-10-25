package main

import (
	"bytes"
	"fmt"
	"github.com/go-ai-agent/core/exchange"
	"github.com/go-ai-agent/core/httpx"
	"github.com/go-ai-agent/core/runtime"
	"net/http"
	"net/url"
)

const (
	ActivityUrl   = "http://localhost:8080/go-ai-agent/example-domain/activity/entry"
	SloUrl        = "http://localhost:8080/go-ai-agent/example-domain/slo/entry"
	TimeseriesUrl = "http://localhost:8080/go-ai-agent/example-domain/timeseries/entry"

	ActivityResource   = "file://[cwd]/pkg/resource/activity.json"
	SloResource        = "file://[cwd]/pkg/resource/slo.json"
	TimeseriesResource = "file://[cwd]/pkg/resource/timeseries.json"
)

func main() {
	Put(ActivityResource, ActivityUrl)
	Put(SloResource, SloUrl)
	Put(TimeseriesResource, TimeseriesUrl)

	//Delete(ActivityUrl)
	//Delete(SloUrl)
	//Delete(TimeseriesUrl)
}

func Put(file, uri string) {
	u, _ := url.Parse(file)
	buf, err := httpx.ReadFile(u)
	if err != nil {
		fmt.Printf("read file err: %v\n", err)
		return
	}
	reader := bytes.NewReader(buf)
	req, err1 := http.NewRequest(http.MethodPut, uri, reader)
	if err1 != nil {
		fmt.Printf("new request err: %v\n", err1)
		return
	}
	resp, _ := exchange.Do[runtime.DebugError](req)
	if resp != nil {
		fmt.Printf("StatusCode: %v\n", resp.StatusCode)
	}
}

func Delete(uri string) {
	req, err1 := http.NewRequest(http.MethodDelete, uri, nil)
	if err1 != nil {
		fmt.Printf("new request err: %v\n", err1)
		return
	}
	resp, _ := exchange.Do[runtime.DebugError](req)
	if resp != nil {
		fmt.Printf("StatusCode: %v\n", resp.StatusCode)
	}
}
