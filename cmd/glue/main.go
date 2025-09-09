package main

import (
	"flag"
	"os"

	"github.com/formalco/rpcglue"
	"github.com/formalco/rpcglue/log"
	"github.com/formalco/rpcglue/provider"
	"github.com/formalco/rpcglue/provider/gorilla"
	"github.com/formalco/rpcglue/provider/stl"
	"github.com/formalco/rpcglue/writer"
)

var debug = flag.Bool("debug", false, "enable debug logs")

// Required
var name = flag.String("name", "", "target RPC declaration name (e.g. Service in `type Service struct`)")
var service = flag.String("service", "", "RPC service name (e.g. `Service` in `Service.Method`)")

// Overrides
var out = flag.String("out", "./client", "output directory")
var print = flag.Bool("print", false, "output code to stdout instead of file")

// Custom providers (only pick one)
var gorillaFlag = flag.Bool("gorilla", false, "supports Gorilla rpc method format")

func main() {
	flag.Parse()

	if *debug {
		log.DebugMode = true
	}

	if *service == "" {
		log.Print("-service is required")
		os.Exit(2)
	}

	if *name == "" {
		log.Print("-name is required")
		os.Exit(2)
	}

	var wr writer.Writer
	if *print {
		wr = writer.NewStdoutWriter()
	} else {
		var err error
		wr, err = writer.NewFileWriter(*out)
		if err != nil {
			os.Exit(1)
		}
	}

	var provider provider.Provider = &stl.Provider{}
	if *gorillaFlag {
		provider = gorilla.New(provider)
	}

	walker := rpcglue.Walker{
		Provider: provider,
		Writer:   wr,
	}

	var path string
	args := flag.Args()
	if len(args) == 0 {
		path = "."
	} else {
		path = args[0]
	}

	err := walker.Walk(rpcglue.Directions{
		Path:    path,
		Name:    *name,
		Service: *service,
	})
	if err != nil {
		os.Exit(2)
	}
}
