package request

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func Send(searchQuery, apiKey, apiSecret string) ([]byte, error) {
	url := fmt.Sprintf("https://api.podcastindex.org/api/1.0/%s", searchQuery)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-Auth-Key", apiKey)
	req.Header.Set("User-Agent", "Podkasto/0.0.1")

	headerTime := strconv.FormatInt(time.Now().Unix(), 10)
	req.Header.Set("X-Auth-Date", headerTime)

	h := sha1.New()
	h.Write([]byte(fmt.Sprintf("%s%s%s", apiKey, apiSecret, headerTime)))
	req.Header.Set("Authorization", fmt.Sprintf("%x", h.Sum(nil)))

	client := http.Client{Timeout: time.Second * 33}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}

	return body, nil
}
