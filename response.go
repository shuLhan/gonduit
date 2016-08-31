// Copyright 2016 Muhammad Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package gonduit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//
// ResponseObject contain object returned by create or edit API.
//
type ResponseObject struct {
	ID   int    `json:"id"`
	Phid string `json:"phid"`
}

//
// ResponseTransaction contain transaction information returned by create or
// edit API.
//
type ResponseTransaction struct {
	Phid string `json:"phid"`
}

//
// ResponseResult contain information returned by create or edit API.
//
type ResponseResult struct {
	Object       ResponseObject        `json:"object"`
	Transactions []ResponseTransaction `json:"transactions"`
}

//
// Response contain information returned by Conduit API (create, edit, or
// search).
//
type Response struct {
	body    []byte
	Result  json.RawMessage `json:"result"`
	ErrCode string          `json:"error_code"`
	ErrInfo string          `json:"error_info"`
}

//
// NewResponse will create new response object for client `cl` and initialize
// it with HTTP response `httpResp` returned by `Request.Post`.
//
func (cl *Client) NewResponse(httpResp *http.Response) (e error) {
	if DEBUG >= 3 {
		fmt.Printf("[gonduit] Response.HTTP << %+v\n", httpResp)
	}

	cl.respon = &Response{}

	cl.respon.body, e = ioutil.ReadAll(httpResp.Body)

	if e != nil {
		return e
	}

	if DEBUG >= 2 {
		fmt.Printf("[gonduit] Response.BODY << %s\n", cl.respon.body)
	}

	e = json.Unmarshal(cl.respon.body, &cl.respon)

	if e != nil {
		return e
	}

	e = httpResp.Body.Close()

	return e
}

//
// DecodeResult will decode response data returned by create/edit API.
//
func (resp *Response) DecodeResult() (rr *ResponseResult, e error) {
	rr = &ResponseResult{}

	e = json.Unmarshal(resp.Result, rr)

	if e != nil {
		return nil, e
	}

	if DEBUG >= 2 {
		fmt.Printf("[gonduit] Response.DecodeResult << %+v\n", rr)
	}

	return rr, nil
}
