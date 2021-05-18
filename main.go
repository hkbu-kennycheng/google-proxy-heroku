package main

import (
    "net/http"
    "net/http/httputil"
    "net/url"
    "os"
)

func main() {
    proxy := &httputil.ReverseProxy{
        Director: func(req *http.Request) {
            u, _ := url.Parse(os.Getenv("UPSTREAM_SERVER"))
	    p := u.Path + req.URL.Path
	    req.URL = u
	    req.URL.Path = p
            req.Host = u.Host
        },
    }
    http.ListenAndServe(":"+os.Getenv("PORT"), proxy)
}
