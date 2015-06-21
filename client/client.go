// Package client for http-works
package client

import (
	"net/http"
)

// based on std.client

const (
	DIRECT = 1 << iota
	WEBDRIVER
	PROXY
	DISTRIBUTED
)

type ProxyPool struct {
}

type ClientProfile struct {
	UserAgent    string
	GzipBububu   string
	CacheTimeout int
	Transort     string // http, https, spdy, http2
}

type Param struct {
}

func Get(s string) (interface{}, interface{}) {
	return http.Get(s)
}

// TODO: Adopt template for mass usage

func worker(linkChan chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for url := range linkChan {
	}
}

func likemain() {
	lCh := make(chan string)
	wg := new(sync.WaitGroup)
	for i := 0; i < 250; i++ {
		wg.Add(1)
		go worker(lCh, wg)
	}
	for _, link := range yourLinksSlice {
		lCh <- link
	}
	close(lCh)
	wg.Wait()
}
