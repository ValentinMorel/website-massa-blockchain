package rpc

import (
	"net/http"

	jsonrpc "github.com/ybbus/jsonrpc/v3"
)

var (
	NETWORK_ADDRESS = "https://inno.massa.net/test13"
)

type Client struct {
	RPCClient jsonrpc.RPCClient
	Url       string
}

func NewClient() *Client {
	client := &Client{
		Url: NETWORK_ADDRESS,
	}
	client.RPCClient =
		jsonrpc.NewClientWithOpts(
			NETWORK_ADDRESS,
			&jsonrpc.RPCClientOpts{
				HTTPClient: &http.Client{},
			})

	return client

}
