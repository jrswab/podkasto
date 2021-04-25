package search

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/jrswab/podkasto/request"
)

func (q *API) ByFeedID(id string) {
	searchQuery := fmt.Sprintf("podcasts/byfeedid?id=%s&pretty", id)

	body, err := request.Send(searchQuery, q.ApiKey, q.ApiSecret)
	if err != nil {
		log.Fatalf("query \"ByFeedID\" failed: %s", err)
	}

	resp := new(Podcasts)
	err = json.Unmarshal(body, &resp)
	if err != nil {
		log.Fatalf("query \"ByFeedID\" failed: %s", err)
	}

	fmt.Printf("%+v\n", resp.Feed)
}
