package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	spinhttp "github.com/fermyon/spin/sdk/go/v2/http"
)

func init() {
	spinhttp.Handle(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintln(w, "This is a silly demo running as WASM!")
		logger := log.New(os.Stderr, "", log.LstdFlags)
		logger.Println("This is a silly demo running as WASM!")
	})
}

func main() {}
