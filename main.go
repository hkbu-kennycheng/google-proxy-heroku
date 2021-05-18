package main

import (
    "net/http"
    "net/http/httputil"
    "net/url"
    "os"
)

func main() {
    u, _ := url.Parse(os.Getenv("UPSTREAM_SERVER"))
    proxy := &httputil.ReverseProxy{
        Director: func(req *http.Request) {
	    u.Path = u.Path + req.URL.Path
	    req.URL = u
            req.Host = u.Host
        },
    }
    http.ListenAndServe(":"+os.Getenv("PORT"), proxy)
}
