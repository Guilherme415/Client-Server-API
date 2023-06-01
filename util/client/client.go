package client

import "net/http"

type IClient interface {
	Do(req *http.Request) (*http.Response, error)
}
