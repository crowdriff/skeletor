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

var (
	bind = flag.CommandLine.String("bind", "127.0.0.1:8000",
		"<address>:<port> to bind HTTP server")
	prod = flag.CommandLine.Bool("prod", false, "Production flag")
)

func main() {
	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Parse the flags and store them as a conf struct
	conf := skeletor.Conf{
		Prod: *prod,
	}
	skeletor.SaveConf(&conf)

	// Start the web server
	graceful.AddSignal(syscall.SIGINT, syscall.SIGTERM)
	app := api.New()
	log.Println("Skeletor server listening on " + *bind)
	err := graceful.ListenAndServe(*bind, app)
	if err != nil {
		log.Fatal(err)
	}

	graceful.Wait()
}
