// Package client for http-works
package client

import (
	"http"
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
