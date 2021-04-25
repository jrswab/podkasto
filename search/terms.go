package search

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/jrswab/podkasto/request"
)

type Podcasts struct {
	Status      string  `json:"status"`
	Feeds       []Feeds `json:"feeds"`
	Count       int     `json:"count"`
	Query       string  `json:"query"`
	Description string  `json:"description"`
}

type Categories struct {
	Num102 string `json:"102"`
}

type Feeds struct {
	ID                     int        `json:"id"`
	Title                  string     `json:"title"`
	URL                    string     `json:"url"`
	Originalurl            string     `json:"originalUrl"`
	Link                   string     `json:"link"`
	Description            string     `json:"description"`
	Author                 string     `json:"author"`
	Ownername              string     `json:"ownerName"`
	Image                  string     `json:"image"`
	Artwork                string     `json:"artwork"`
	Lastupdatetime         int        `json:"lastUpdateTime"`
	Lastcrawltime          int        `json:"lastCrawlTime"`
	Lastparsetime          int        `json:"lastParseTime"`
	Lastgoodhttpstatustime int        `json:"lastGoodHttpStatusTime"`
	Lasthttpstatus         int        `json:"lastHttpStatus"`
	Contenttype            string     `json:"contentType"`
	Itunesid               int        `json:"itunesId"`
	Generator              string     `json:"generator"`
	Language               string     `json:"language"`
	Type                   int        `json:"type"`
	Dead                   int        `json:"dead"`
	Crawlerrors            int        `json:"crawlErrors"`
	Parseerrors            int        `json:"parseErrors"`
	Categories             Categories `json:"categories"`
	Locked                 int        `json:"locked"`
	Imageurlhash           int        `json:"imageUrlHash"`
}

type Query struct {
	ApiKey    string
	ApiSecret string
}

func (q *Query) ByTerm(term string) {
	//searchQuery := fmt.Sprintf("%sbyterm?q=%s&pretty=%s", searchBase, query, isPretty)
	searchQuery := fmt.Sprintf("byterm?q=%s", term)

	body, err := request.Send(searchQuery, q.ApiKey, q.ApiSecret)
	if err != nil {
		log.Fatal(err)
	}

	search := new(Podcasts)
	err = json.Unmarshal(body, &search)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("%+v", search)
}
