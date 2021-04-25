package main

import (
	"flag"

	"github.com/jrswab/podkasto/config"
	"github.com/jrswab/podkasto/search"
)

func main() {
	var term = flag.String("term", "", "A term to search for.")
	flag.Parse()

	conf := new(config.Config)
	conf.Load()

	if *term != "" {
		q := &search.Query{ApiKey: conf.ApiKey, ApiSecret: conf.ApiSecret}
		q.ByTerm(*term)
	}
}
