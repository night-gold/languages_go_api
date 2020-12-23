package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/valyala/fasthttp"
)

var (
	addr = flag.String("addr", ":8080", "TCP address to listen to, default to *:8080")
)

func main() {
	flag.Parse()

	h := requestHandler

	if err := fasthttp.ListenAndServe(*addr, h); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Hello, world!\n\n")
}
