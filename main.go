package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"

	"comail.io/go/colog"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	config  = kingpin.Flag("config", "configuration file to create endpoints from").Short('c').Default("./settings.json").String()
	verbose = kingpin.Flag("verbose", "enabled verbose logging. if --silent is used, --verbose will be ignore").Short('v').Default("false").Bool()
	silent  = kingpin.Flag("silent", "disabled all output except for errors. overrides --verbose if set").Short('s').Default("false").Bool()
)

func main() {
	kingpin.Parse()
	initLogging()

	f, err := os.Open(*config)
	if err != nil {
		log.Fatal(err)
	}
	settings, err := NewSettings(f)
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	setupHandlers(settings, r)

	srv := &http.Server{Addr: fmt.Sprintf(":%d", settings.Port()), Handler: r}
	srv.RegisterOnShutdown(func() {
		log.Printf("info: shutting down server")
	})

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		for range c {
			srv.Shutdown(context.Background())
		}
	}()
	log.Printf("info: starting server at %s", srv.Addr)
	srv.ListenAndServe()
}

func setupHandlers(s *Settings, r *mux.Router) {
	for _, m := range s.Mocks {
		f := createHandler(m.ResponseCode, m.ResponseBody, m.Headers)
		r.HandleFunc(m.Path, f).Methods(m.Method)
		log.Printf("debug: creating mock for %s %s - returning %d, %s using %v", m.Method, m.Path, m.ResponseCode, m.ResponseBody, f)
	}
}

func initLogging() {
	colog.Register()
	colog.SetDefaultLevel(colog.LDebug)
	if *silent {
		colog.SetMinLevel(colog.LError)
		return
	}

	if *verbose {
		colog.SetMinLevel(colog.LDebug)
	} else {
		colog.SetMinLevel(colog.LInfo)
	}
	log.Printf("debug: enable verbose logging")
}
