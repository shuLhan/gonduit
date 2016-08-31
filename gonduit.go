// Copyright 2016 Muhammad Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package gonduit

import (
	"os"
	"strconv"
)

var (
	// DEBUG debug level, set using environment GONDUIT_DEBUG
	DEBUG = 0
)

func init() {
	var e error
	DEBUG, e = strconv.Atoi(os.Getenv("GONDUIT_DEBUG"))
	if e != nil {
		DEBUG = 0
	}
}
