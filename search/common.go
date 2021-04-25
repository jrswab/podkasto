package search

// API contains the API Key and Secret for sending requests to PodcastIndex.org
type API struct {
	ApiKey    string
	ApiSecret string
}

// Response contains the unmarshaled data from a successful query to PodcastIndex.org
type Search struct {
	Status      string `json:"status"`
	Query       string `json:"query"`
	Feeds       []Feed `json:"feeds"`
	Description string `json:"description"`
	Count       int    `json:"count"`
}

type Podcasts struct {
	Status      string `json:"status"`
	Query       Query  `json:"query"`
	Feed        Feed   `json:"feed"`
	Description string `json:"description"`
	Count       int    `json:"count"`
}

type Query struct {
	ID string `json:"id"`
}

// Feed contains the unmarshaled data about a returned feed from PodcastingIndex.org
type Feed struct {
	ID                     int    `json:"id"`
	Title                  string `json:"title"`
	URL                    string `json:"url"`
	Originalurl            string `json:"originalUrl"`
	Link                   string `json:"link"`
	Description            string `json:"description"`
	Author                 string `json:"author"`
	Ownername              string `json:"ownerName"`
	Image                  string `json:"image"`
	Artwork                string `json:"artwork"`
	Lastupdatetime         int    `json:"lastUpdateTime"`
	Lastcrawltime          int    `json:"lastCrawlTime"`
	Lastparsetime          int    `json:"lastParseTime"`
	Lastgoodhttpstatustime int    `json:"lastGoodHttpStatusTime"`
	Lasthttpstatus         int    `json:"lastHttpStatus"`
	Contenttype            string `json:"contentType"`
	Itunesid               int    `json:"itunesId"`
	Generator              string `json:"generator"`
	Language               string `json:"language"`
	Type                   int    `json:"type"`
	Dead                   int    `json:"dead"`
	Chash                  string `json:"chash"`
	EpisodeCount           int    `json:"episodeCount"`
	Crawlerrors            int    `json:"crawlErrors"`
	Parseerrors            int    `json:"parseErrors"`
	Locked                 int    `json:"locked"`
	Imageurlhash           int    `json:"imageUrlHash"`
}
