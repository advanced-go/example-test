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

func main() {
	LoadActivity()
}

func LoadActivity() {
	s := "file://[cwd]/pkg/resource/activity.json"
	u, _ := url.Parse(s)

	buf, err := httpx.ReadFile(u)
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}
	//fmt.Printf("buf: %v", string(buf))

	uri := "http://localhost:8080/go-ai-agent/example-domain/activity/entry"
	reader := bytes.NewReader(buf)
	req, err1 := http.NewRequest(http.MethodPut, uri, reader)
	if err1 != nil {
		fmt.Printf("err: %v", err1)
	}

	resp, _ := exchange.Do[runtime.DebugError](req)
	if resp != nil {
		fmt.Printf("StatusCode: %v", resp.StatusCode)
	}

}
