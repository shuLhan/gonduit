// Copyright 2016 Muhammad Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package gonduit

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

//
// Request contain API and parameter to be sent to Conduit server.
//
type Request struct {
	api  string
	v    url.Values
	body strings.Reader

	idx int
	str string
}

//
// NewRequest create new request values for client.
//
func (cl *Client) NewRequest(api string) {
	cl.request = &Request{
		api: cl.url + api,
		v:   url.Values{},
		idx: 0,
		str: "",
	}

	cl.request.v.Set(APIParamToken, cl.token)
}

//
// EncodeValues will transform request parameters into string.
//
func (req *Request) EncodeValues() string {
	return req.v.Encode()
}

//
// AddParam will add parameter to be sent on request object.
//
func (req *Request) AddParam(t, v string) {
	req.str = fmt.Sprintf("transactions[%d][type]", req.idx)
	req.v.Add(req.str, t)

	req.str = fmt.Sprintf("transactions[%d][value]", req.idx)
	req.v.Add(req.str, v)

	req.idx++
}

//
// AddParams will add list of parameter to be sent on request object.
//
func (req *Request) AddParams(t string, vals []string) {
	req.str = fmt.Sprintf("transactions[%d][type]", req.idx)
	req.v.Add(req.str, t)

	for x, v := range vals {
		req.str = fmt.Sprintf("transactions[%d][value][%d]",
			req.idx, x)

		req.v.Add(req.str, v)
	}
}

//
// AddConstraint will add parameter constraint for API required in search.
//
func (req *Request) AddConstraint(name, value string) {
	req.str = fmt.Sprintf("constraints[%s]", name)
	req.v.Add(req.str, value)
}

//
// Post sent the request to the server.
//
func (req *Request) Post(cl *Client) (*http.Response, error) {
	var httpReq *http.Request
	var e error

	fmt.Printf("REQ  >> %v\n", req.api)
	fmt.Printf("BODY >> %v\n", req.v.Encode())

	req.body.Reset(req.v.Encode())

	httpReq, e = http.NewRequest("POST", req.api, &req.body)

	if e != nil {
		return nil, e
	}

	return cl.http.Do(httpReq)
}
