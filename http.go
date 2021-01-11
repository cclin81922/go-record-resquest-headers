package main

import (
    "fmt"
    "net/http"
)

func headers(w http.ResponseWriter, req *http.Request) {
    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Printf("%v: %v\n", name, h)
        }
    }
}

func HttpServe() {
    http.HandleFunc("/", headers)
    go http.ListenAndServe(":8080", nil)
}
