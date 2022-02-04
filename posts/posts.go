package posts

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Posts struct {
	Client *http.Client
	Url    string // /v1/posts/?auth_token=...
}

type Filter struct {
	Public bool

	UI         *string
	Currencies *string
	Regions    *string
	Kind       *string
}

type LookupResults struct {
	Kind   string `json:"kind"`
	Domain string `json:"domain"`
	Votes  struct {
		Negative  int `json:"negative"`
		Positive  int `json:"positive"`
		Important int `json:"important"`
		Liked     int `json:"liked"`
		Disliked  int `json:"disliked"`
		Lol       int `json:"lol"`
		Toxic     int `json:"toxic"`
		Saved     int `json:"saved"`
		Comments  int `json:"comments"`
	} `json:"votes"`
	Source struct {
		Title  string      `json:"title"`
		Region string      `json:"region"`
		Domain string      `json:"domain"`
		Path   interface{} `json:"path"`
	} `json:"source"`
	Title       string    `json:"title"`
	PublishedAt time.Time `json:"published_at"`
	Slug        string    `json:"slug"`
	Currencies  []struct {
		Code  string `json:"code"`
		Title string `json:"title"`
		Slug  string `json:"slug"`
		URL   string `json:"url"`
	} `json:"currencies,omitempty"`
	ID        int       `json:"id"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}

type Lookup struct {
	Count    int             `json:"count"`
	Next     string          `json:"next"`
	Previous interface{}     `json:"previous"`
	Results  []LookupResults `json:"results"`
}

func (at *Posts) Latest(rows int, filter Filter) []LookupResults {
	url := at.Url
	if filter.Public {
		url += "&public=true"
	}
	if filter.UI != nil {
		url += "&filter=" + *filter.UI
	}
	if filter.Currencies != nil {
		url += "&currencies=" + *filter.Currencies
	}
	if filter.Regions != nil {
		url += "&regions=" + *filter.Regions
	}
	if filter.Kind != nil {
		url += "&kind=" + *filter.Kind
	}

	req, reqEx := http.NewRequest(http.MethodGet, url, nil)
	if reqEx != nil {
		log.Fatal(reqEx)
	}

	res, resEx := at.Client.Do(req)
	if resEx != nil {
		log.Fatal(resEx)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readEx := ioutil.ReadAll(res.Body)
	if readEx != nil {
		log.Fatal(readEx)
	}

	lookup := Lookup{}
	jsonEx := json.Unmarshal(body, &lookup)

	if jsonEx != nil {
		log.Fatal(jsonEx)
	}

	return lookup.Results
}
