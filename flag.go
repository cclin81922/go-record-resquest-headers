package main

import (
    "github.com/namsral/flag"
)

var (
    port = flag.Int("port", 8080, "HTTP server port")
)

func parseArgs() {
    flag.Parse()
}
