package main

import (
	"flag"
	"github.com/daniyalibrahim/go-fibonacci/internal/fibonacciapi"
	"github.com/daniyalibrahim/go-fibonacci/pkg/log"
	"net/http"
)

// Version indicates the current version of the application.
var Version = "1.0.0"

var flagConfig = flag.String("config", "../config/local.yml", "path to the config file")

func main() {
	flag.Parse()
	// create root logger tagged with server version
	logger := log.New().With(nil, "version", Version)

	FibonacciHandlers := fibonacciapi.NewFibonacciHandlers()
	http.HandleFunc("/helloapi/fibonacci", FibonacciHandlers.Get)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		logger.Error(err)
		panic(err)
	}
}
