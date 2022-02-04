package cryptopanic

// https://cryptopanic.com/developers/api/

import (
	"github.com/jerskisnow/CryptoPanic-GO/posts"
	"net/http"
)

type Client struct {
	client   *http.Client
	baseURL  string
	apiToken string
}

func CreateClient(token string) *Client {
	return &Client{
		client:   http.DefaultClient,
		baseURL:  "https://cryptopanic.com/api/v1/",
		apiToken: token,
	}
}

func (at *Client) Posts() *posts.Posts {
	return &posts.Posts{
		Client: at.client,
		Url:    at.baseURL + "posts/?auth_token=" + at.apiToken,
	}
}
