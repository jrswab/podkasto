package search

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/jrswab/podkasto/request"
)

func (q *API) ByTerm(term string) {
	//searchQuery := fmt.Sprintf("%sbyterm?q=%s&pretty=%s", searchBase, query, isPretty)
	searchQuery := fmt.Sprintf("search/byterm?q=%s", term)

	body, err := request.Send(searchQuery, q.ApiKey, q.ApiSecret)
	if err != nil {
		log.Fatal(err)
	}

	response := new(Search)
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatalf("query \"ByTerm\" failed: %s", err)
	}

	for _, feed := range response.Feeds {
		fmt.Printf("%+v\n\n", feed)
	}
}

func (q *API) ByPerson(name string) {
	searchQuery := fmt.Sprintf("search/byperson?q=%s", name)

	body, err := request.Send(searchQuery, q.ApiKey, q.ApiSecret)
	if err != nil {
		log.Fatal(err)
	}

	response := new(Search)
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatalf("query \"ByPerson\" failed: %s", err)
	}

	for _, feed := range response.Feeds {
		fmt.Printf("%+v\n\n", feed)
	}
}
