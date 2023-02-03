package utils

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	headerContentEncoding = "Content-Encoding"
	encodingGzip          = "gzip"
)

type ProxyOpt struct {
	Addr string `yaml:"host"`
}

func DoProxy(w http.ResponseWriter, r *http.Request, opt ProxyOpt) {
	cli := &http.Client{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		zap.S().Warnf("err to read request body: %v", err)
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	reqURL := opt.Addr + r.URL.String()
	proxyRequest, err := http.NewRequest(r.Method, reqURL, strings.NewReader(string(body)))
	if err != nil {
		zap.S().Warnf("err to create proxy request: %v", err)
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	for k, v := range r.Header {
		proxyRequest.Header.Set(k, v[0])
	}
	proxyResponse, err := cli.Do(proxyRequest)
	if err != nil {
		zap.S().Warnf("err to transmit request: %v", err)
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	defer proxyResponse.Body.Close()
	for k, v := range proxyResponse.Header {
		w.Header().Set(k, v[0])
	}
	var data []byte
	data, err = ioutil.ReadAll(proxyResponse.Body)
	if err != nil {
		zap.S().Warnf("err to read response body: %v", err)
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	var dataOutput []byte
	isGzipped := isGzipped(proxyResponse.Header)
	if isGzipped {
		// 读取后 r.Body 即关闭，无法再次读取
		// 若需要再次读取，需要用读取到的内容再次构建Reader
		resProxyGzippedBody := ioutil.NopCloser(bytes.NewBuffer(data))
		defer resProxyGzippedBody.Close()
		gzipReader, err := gzip.NewReader(resProxyGzippedBody)
		if err != nil {
			zap.S().Warnf("err to create gzip reader: %v", err)
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		defer gzipReader.Close()
		dataOutput, err = ioutil.ReadAll(gzipReader)
		if err != nil {
			zap.S().Warnf("err to read gzip: %v", err)
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
	} else {
		dataOutput = data
	}

	println(string(dataOutput))

	// response的Body不能多次读取，
	// 上面已经被读取过一次，需要重新生成可读取的Body数据。
	resProxyBody := ioutil.NopCloser(bytes.NewBuffer(data))
	defer resProxyBody.Close()
	w.WriteHeader(proxyResponse.StatusCode)
	io.Copy(w, resProxyBody)
}

func isGzipped(header http.Header) bool {
	if header == nil {
		return false
	}
	contentEncoding := header.Get(headerContentEncoding)
	isGzipped := false
	if strings.Contains(contentEncoding, encodingGzip) {
		isGzipped = true
	}
	return isGzipped
}

func Proxy(c *gin.Context) {
	// fmt.Println(c.GetHeader("direct"))
	// if c.GetHeader("direct") != "lab" {
	// 	return
	// }
	err := setTokenToUrl(c.Request.URL)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("addr err: %s", err.Error()))
		c.Abort()
		return
	}
	req, err := http.NewRequestWithContext(c, c.Request.Method, c.Request.URL.String(), c.Request.Body)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		c.Abort()
		return
	}
	defer req.Body.Close()
	req.Header = c.Request.Header
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		c.Abort()
		return
	}
	for k := range resp.Header {
		for j := range resp.Header[k] {
			c.Header(k, resp.Header[k][j])
		}
	}
	extraHeaders := make(map[string]string)
	extraHeaders["direct"] = "lab"
	c.DataFromReader(resp.StatusCode, resp.ContentLength, resp.Header.Get("Content-Type"), resp.Body, extraHeaders)
	c.Abort()
}

func setTokenToUrl(rawUrl *url.URL) error {
	proxyUrl := "http://127.0.0.1:9095/"
	// token := ""
	u, err := url.Parse(proxyUrl)
	if err != nil {
		return err
	}
	rawUrl.Scheme = u.Scheme
	rawUrl.Host = u.Host
	ruq := rawUrl.Query()
	// ruq.Add("token", token)
	rawUrl.RawQuery = ruq.Encode()
	return nil
}
