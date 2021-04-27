package main

import (
	"flag"

	"github.com/jrswab/podkasto-api/search"
	"github.com/jrswab/podkasto-cli/config"
)

func main() {
	var (
		searchType = flag.String("search", "",
			"May be set to \"podcasts\" or \"episodes\".\nIf left unset, podkasto will perform a basic query to PodcastIndex.org")
		term = flag.String("term", "",
			"Used to search for podcasts containing a term.\nExample: podkasto --search podcasts --term \"Hacker Culture\"")
		id = flag.String("id", "",
			"Used to search for a podcast by the \"Podcast Index\" id.\nExample: podkasto --search podcasts --id 669240")
		name = flag.String("name", "",
			"Search for podcasts by author name. Do not include the \"search\" flag with this one.\nExpample: podkasto --name \"Jaron Swab\"")
	)
	flag.Parse()

	conf := new(config.Config)
	conf.Load()

	q := &search.API{ApiKey: conf.ApiKey, ApiSecret: conf.ApiSecret}
	if *searchType == "podcasts" {
		if *term != "" {
			q.ByTerm(*term)
		}

		if *id != "" {
			q.ByFeedID(*id)
		}
	} else {
		if *term != "" {
			q.ByTerm(*term)
		} else if *name != "" {
			q.ByPerson(*name)
		} else {
			flag.PrintDefaults()
		}
	}
}
