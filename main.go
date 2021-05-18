package main

import (
    "net/http"
    "net/http/httputil"
    "net/url"
    "os"
)

func main() {
    u, _ := url.Parse(os.Getenv("UPSTREAM_SERVER"))
    http.Handle("/", httputil.NewSingleHostReverseProxy(u))
 
    // Start the server
    http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
