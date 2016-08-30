// Copyright 2016 Muhammad Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package gonduit_test

import (
	"github.com/shuLhan/gonduit"
	"testing"
)

const (
	CONDUIT_URL   = "http://phabricator.localdomain:10010/"
	CONDUIT_TOKEN = "api-pcbhtiluxuarqcvg44a4vfkyikw5"
)

func TestPing(t *testing.T) {
	cl := gonduit.NewClient(CONDUIT_URL, CONDUIT_TOKEN)

	e := cl.Ping()

	if e != nil {
		t.Fatal(e)
	}
}
