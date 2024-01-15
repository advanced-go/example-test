package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/advanced-go/core/exchange"
	"github.com/advanced-go/core/runtime"
	slo "github.com/advanced-go/example-domain/slo"
	"io"
	"net/http"
)

const (
	ActivityUrl    = "http://localhost:8080/github/advanced-go/example-domain/service:activity/entry"
	SloUrl         = "http://localhost:8080/github/advanced-go/example-domain/service:slo/entry"
	Timeseries1Url = "http://localhost:8080/github/advanced-go/example-domain/service:timeseries/v1/entry"
	Timeseries2Url = "http://localhost:8080/github/advanced-go/example-domain/service:timeseries/v2/entry"

	ActivityResource     = "file://[cwd]/pkg/resource/activity.json"
	SloResource          = "file://[cwd]/pkg/resource/slo.json"
	TimeseriesResourceV1 = "file://[cwd]/pkg/resource/timeseries-v1.json"
	//TimeseriesResourceV2 = "file://[cwd]/pkg/resource/timeseries-v2.json"
	TimeseriesResourceV2 = "file://[cwd]/pkg/resource/timeseries-v2-annotated.json"
)

func main() {
	//testInitialLoad()
	testAgentLoad()

	//testAgentAddSLO("103", "host", "99.9/701ms")
	//testAgentAddSLO("104", "host", "99.9/801ms")
	//testAgentAddSLO("105", "host", "99.9/901ms")
	//testAgentAddSLO("106", "host", "99.9/1001ms")

	//Delete(ActivityUrl)
	//Delete(SloUrl)
	//Delete(TimeseriesUrl)
}

func testInitialLoad() {
	Put(ActivityResource, ActivityUrl, "")
	Put(SloResource, SloUrl, "")
	Put(TimeseriesResourceV1, Timeseries1Url, "")
	Put(TimeseriesResourceV2, Timeseries2Url, "") //timeseriesvar.EntryV2Variant)

}

func testAgentLoad() bool {
	if !Put(SloResource, SloUrl, "") {
		return false
	}
	return Put(TimeseriesResourceV2, Timeseries2Url, "")
}

func testAgentAddSLO(id, controller, threshold string) bool {
	entries := []slo.EntryV1{{
		Id:          id,
		Controller:  controller,
		Threshold:   threshold,
		StatusCodes: "0",
	},
	}
	buf, err := json.Marshal(entries)
	if err != nil {
		fmt.Printf("error: AddSLO() -> %v", err)
		return false
	}
	r := bytes.NewReader(buf)
	req, err1 := http.NewRequest(http.MethodPut, SloUrl, io.NopCloser(r))
	if err1 != nil {
		fmt.Printf("new request err: %v\n", err1)
		return false
	}
	resp, _ := exchange.Do(req)
	if resp != nil {
		fmt.Printf("StatusCode: %v\n", resp.StatusCode)
	}
	return true
}

func Put(file, uri, variant string) bool {
	//u, _ := url.Parse(file)
	buf, status := runtime.NewBytes(file) //io2.ReadFile(u)
	if !status.OK() {
		fmt.Printf("read file err: %v\n", status.Errors())
		return false
	}
	reader := bytes.NewReader(buf)
	req, err1 := http.NewRequest(http.MethodPut, uri, reader)
	if err1 != nil {
		fmt.Printf("new request err: %v\n", err1)
		return false
	}
	resp, status := exchange.Do(req)
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
	resp, _ := exchange.Do(req)
	if resp != nil {
		fmt.Printf("StatusCode: %v\n", resp.StatusCode)
	}
}
