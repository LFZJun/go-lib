package proxy

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/elazarl/goproxy"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var Port = ":8080"

func init() {
	go Show()
	os.Setenv("http_proxy", "http://127.0.0.1"+Port)
}

func Show() {
	var server = goproxy.NewProxyHttpServer()
	const sp = "------\n"
	server.Tr.Proxy = nil
	server.Tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	server.OnRequest().DoFunc(func(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
		var b strings.Builder
		b.WriteString(sp)
		// first line
		b.WriteString(fmt.Sprintf("%v %v %v\n", req.Method, req.URL.Path, req.Proto))
		// HOST
		b.WriteString(fmt.Sprintf("HOST: %v\n", req.Host))
		// header
		for k, v := range req.Header {
			for _, vv := range v {
				b.WriteString(fmt.Sprintf("%v: %v\n", k, vv))
			}
		}
		b.WriteString("\n\n")
		// body
		if req.Body == http.NoBody {
			b.WriteString(sp)
			fmt.Println(b.String())
			return req, nil
		}
		var buf bytes.Buffer
		_, err := buf.ReadFrom(req.Body)
		if err != nil {
			b.WriteString("---error---\n")
			b.WriteString(err.Error() + "\n")
			b.WriteString(sp)
			fmt.Println(b.String())
			return req, nil
		}
		body := buf.Bytes()
		b.Write(body)
		b.WriteByte('\n')
		b.WriteString(sp)
		fmt.Println(b.String())
		req.Body = ioutil.NopCloser(bytes.NewReader(body))
		return req, nil
	})
	err := http.ListenAndServe(Port, server)
	if err != nil {
		fmt.Println(err)
		return
	}
}
