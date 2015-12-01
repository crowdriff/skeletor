package main

import (
	"flag"
	"log"
	"runtime"
	"syscall"

	"github.com/crowdriff/skeletor/api"
	"github.com/crowdriff/skeletor/skeletor"
	"github.com/zenazn/goji/graceful"
)

// command line flags parsed at runtime
var (
	bind = flag.CommandLine.String("bind", "127.0.0.1:8000",
		"<address>:<port> to bind HTTP server")
	prod = flag.CommandLine.Bool("prod", false, "Production flag")
)

// instantiated at build time
var version string

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Parse the flags and store them as a conf struct
	flag.Parse()
	conf := skeletor.Conf{
		Prod:    *prod,
		Version: version,
	}
	skeletor.SaveConf(&conf)

	// Start the web server
	graceful.AddSignal(syscall.SIGINT, syscall.SIGTERM)
	app := api.New()
	log.Println("Skeletor server listening on", *bind)
	err := graceful.ListenAndServe(*bind, app)
	if err != nil {
		log.Fatal(err)
	}

	graceful.Wait()
}
