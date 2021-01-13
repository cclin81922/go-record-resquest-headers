package main

import (
    "github.com/cclin81922/go-3rd-party/go-flag"
)

var (
    port = flag.Int("port", 8080, "HTTP server port")
)

func parseArgs() {
    flag.Parse()
}
