// Copyright 2016 Muhammad Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package gonduit

import (
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
// NewClient will create new client for connecting conduit API.
//
func NewClient(URL, token string) *Client {
	cl := &Client{
		http:    nil,
		url:     URL,
		token:   token,
		request: nil,
		respon:  nil,
	}

	cl.http = &http.Client{}

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