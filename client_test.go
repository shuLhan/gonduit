// Copyright 2016 Muhammad Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package gonduit_test

import (
	"testing"
)

func TestPing(t *testing.T) {
	e := cl.Ping()

	if e != nil {
		t.Fatal(e)
	}
}
