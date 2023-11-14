package main

import (
	"bytes"
	"fmt"
	"github.com/go-ai-agent/core/http2"
	"github.com/go-ai-agent/core/io2"
	"github.com/go-ai-agent/example-domain/slo"
	"net/http"
	"net/url"
)

const (
	ActivityUrl   = "http://localhost:8080/go-ai-agent/example-domain/activity/entry"
	SloUrl        = "http://localhost:8080/go-ai-agent/example-domain/slo/entry"
	TimeseriesUrl = "http://localhost:8080/go-ai-agent/example-domain/timeseries/entry"

	ActivityResource     = "file://[cwd]/pkg/resource/activity.json"
	SloResource          = "file://[cwd]/pkg/resource/slo.json"
	TimeseriesResourceV1 = "file://[cwd]/pkg/resource/timeseries-v1.json"
	TimeseriesResourceV2 = "file://[cwd]/pkg/resource/timeseries-v2.json"

	EntryV1Variant = "github.com/go-ai-agent/example-domain/timeseries/EntryV1"
	EntryV2Variant = "github.com/go-ai-agent/example-domain/timeseries/EntryV2"
)

func main() {
	testInitialLoad()
	//testAgent_Load()
	//Delete(ActivityUrl)
	//Delete(SloUrl)
	//Delete(TimeseriesUrl)
}

func testInitialLoad() {
	Put(ActivityResource, ActivityUrl, "")
	Put(SloResource, SloUrl, "")
	Put(TimeseriesResourceV1, TimeseriesUrl, EntryV1Variant)
	Put(TimeseriesResourceV2, TimeseriesUrl, EntryV2Variant)

}

func testAgent_Load() {
	Put(SloResource, SloUrl, "")
	Put(TimeseriesResourceV2, TimeseriesUrl, EntryV2Variant)
}

func testAgent_AddSLO(slo slo.EntryV1) {
	variant := ""
	req, err1 := http.NewRequest(http.MethodPut, SloUrl, nil)
	if err1 != nil {
		fmt.Printf("new request err: %v\n", err1)
		return
	}
	if len(variant) > 0 {
		req.Header.Add(http2.ContentLocation, variant)
	}
	resp, _ := http2.Do(req)
	if resp != nil {
		fmt.Printf("StatusCode: %v\n", resp.StatusCode)
	}

	Put(SloResource, SloUrl, "")

}

func Put(file, uri, variant string) {
	u, _ := url.Parse(file)
	buf, err := io2.ReadFile(u)
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
	if len(variant) > 0 {
		req.Header.Add(http2.ContentLocation, variant)
	}
	resp, status := http2.Do(req)
	if resp != nil {
		fmt.Printf("StatusCode: %v\n", resp.StatusCode)
	}
	fmt.Printf("Put() [status:%v]\n", status)
}

func Delete(uri, variant string) {
	req, err1 := http.NewRequest(http.MethodDelete, uri, nil)
	if err1 != nil {
		fmt.Printf("new request err: %v\n", err1)
		return
	}
	if len(variant) > 0 {
		req.Header.Add(http2.ContentLocation, variant)
	}
	resp, _ := http2.Do(req)
	if resp != nil {
		fmt.Printf("StatusCode: %v\n", resp.StatusCode)
	}
}
