// Copyright 2016 Muhammad Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package gonduit

import (
	"crypto/tls"
	"encoding/json"
	"net/http"
)

//
// Client contains all information for HTTP client to connect to conduti server,
// including reusable request and response object.
//
type Client struct {
	http    *http.Client
	url     string
	token   string
	request *Request
	respon  *Response
}

//
// Cursor contains the value to limit search or information returned by search.
//
type Cursor struct {
	Limit  int    `json:"limit"`
	After  string `json:"after"`
	Before string `json:"before"`
	Order  string `json:"order"`
}

//
// SearchResult contain generic result from querying to server.
//
type SearchResult struct {
	Data   json.RawMessage `json:"data"`
	Maps   json.RawMessage `json:"maps"`
	Query  json.RawMessage `json:"query"`
	Cursor Cursor          `json:"cursor"`
}

//
// NewClient will create new HTTP client for connecting to conduit API server at
// `URL` using user `token`.
//
func NewClient(URL, token string, disableTls bool) *Client {
	cl := &Client{
		http:    nil,
		url:     URL,
		token:   token,
		request: nil,
		respon:  nil,
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: disableTls,
		},
	}

	cl.http = &http.Client{
		Transport: tr,
	}

	return cl
}

//
// Post will sent gonduit Request and parse their response.
//
func (cl *Client) Post() (e error) {
	var httpResp *http.Response

	httpResp, e = cl.request.Post(cl)

	if e != nil {
		return e
	}

	e = cl.NewResponse(httpResp)

	return e
}

//
// Ping check if server is up.
//
func (cl *Client) Ping() error {
	cl.NewRequest(APIPing)

	return cl.Post()
}
