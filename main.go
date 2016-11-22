package main

import (
	"github.com/msawangwan/mlogaccess/maccess"
	"net/http"
)

const (
	LISTEN_PORT       = ":8080"
	PROXY_PASS        = "http://127.0.0.1:8000"
	WHITELIST         = `^\/$|[\w|/]*.js|/path|/path2`
	LISTEN_PORT_USAGE = "listen for requests on this port ... i.e, 80, 8080"
	PROXY_PASS_USAGE  = "redirect target address ... i.e, http://127.0.0.1:8080"
	WHITELIST_USAGE   = "match on whitelisted characters ... i.e, regex"
)

func main() {
	proxyGateway := mlogaccess.NewReverseProxyGateway(PROXY_PASS, WHITELIST)
	proxyGateway.GatewayLogger.LogStatus("redirecting and logging traffic ...")

	http.HandleFunc("/", proxyGateway.ProxyPassHandler)
	http.ListenAndServe(LISTEN_PORT, nil)
}
