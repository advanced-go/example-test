package main

import (
	"bytes"
	"fmt"
	"github.com/advanced-go/core/http2"
	"github.com/advanced-go/core/io2"
	"github.com/advanced-go/core/json2"
	"github.com/advanced-go/example-domain/activity"
	"github.com/advanced-go/example-domain/slo"
	"github.com/advanced-go/example-domain/timeseries"
	"io"
	"net/http"
	"net/url"
)

const (
	ActivityUrl   = "http://localhost:8080/advanced-go/example-domain/activity/entry"
	SloUrl        = "http://localhost:8080/advanced-go/example-domain/slo/entry"
	TimeseriesUrl = "http://localhost:8080/advanced-go/example-domain/timeseries/entry"

	ActivityResource     = "file://[cwd]/pkg/resource/activity.json"
	SloResource          = "file://[cwd]/pkg/resource/slo.json"
	TimeseriesResourceV1 = "file://[cwd]/pkg/resource/timeseries-v1.json"
	TimeseriesResourceV2 = "file://[cwd]/pkg/resource/timeseries-v2.json"
)

func main() {
	//testInitialLoad()
	testAgentLoad()

	//testAgentAddSLO("103","host","99.9/701ms")
	//testAgentAddSLO("104","host","99.9/801ms")
	//testAgentAddSLO("105","host","99.9/901ms")
	//testAgentAddSLO("106","host","99.9/1001ms")

	//Delete(ActivityUrl)
	//Delete(SloUrl)
	//Delete(TimeseriesUrl)
}

func testInitialLoad() {
	Put(ActivityResource, ActivityUrl, activity.EntryV1Variant)
	Put(SloResource, SloUrl, slo.EntryV1Variant)
	Put(TimeseriesResourceV1, TimeseriesUrl, timeseries.EntryV1Variant)
	Put(TimeseriesResourceV2, TimeseriesUrl, timeseries.EntryV2Variant)

}

func testAgentLoad() bool {
	if !Put(SloResource, SloUrl, slo.EntryV1Variant) {
		return false
	}
	return Put(TimeseriesResourceV2, TimeseriesUrl, timeseries.EntryV2Variant)
}

func testAgentAddSLO(id, controller, threshold string) bool {
	entries := []slo.EntryV1{{
		Id:          id,
		Controller:  controller,
		Threshold:   threshold,
		StatusCodes: "0",
	},
	}
	buf, status := json2.Marshal(entries)
	if !status.OK() {
		fmt.Printf("error: AddSLO() -> %v", status.FirstError())
		return false
	}
	r := bytes.NewReader(buf)
	req, err1 := http.NewRequest(http.MethodPut, SloUrl, io.NopCloser(r))
	if err1 != nil {
		fmt.Printf("new request err: %v\n", err1)
		return false
	}
	req.Header.Add(http2.ContentLocation, slo.EntryV1Variant)
	resp, _ := http2.Do(req)
	if resp != nil {
		fmt.Printf("StatusCode: %v\n", resp.StatusCode)
	}
	return true
}

func Put(file, uri, variant string) bool {
	u, _ := url.Parse(file)
	buf, err := io2.ReadFile(u)
	if err != nil {
		fmt.Printf("read file err: %v\n", err)
		return false
	}
	reader := bytes.NewReader(buf)
	req, err1 := http.NewRequest(http.MethodPut, uri, reader)
	if err1 != nil {
		fmt.Printf("new request err: %v\n", err1)
		return false
	}
	req.Header.Add(http2.ContentLocation, variant)
	resp, status := http2.Do(req)
	if resp != nil {
		fmt.Printf("StatusCode: %v\n", resp.StatusCode)
	}
	fmt.Printf("Put() [status:%v]\n", status)
	return true
}

func Delete(uri, variant string) {
	req, err1 := http.NewRequest(http.MethodDelete, uri, nil)
	if err1 != nil {
		fmt.Printf("new request err: %v\n", err1)
		return
	}
	req.Header.Add(http2.ContentLocation, variant)
	resp, _ := http2.Do(req)
	if resp != nil {
		fmt.Printf("StatusCode: %v\n", resp.StatusCode)
	}
}
