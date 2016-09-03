// Copyright 2016 Muhammad Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package gonduit_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/shuLhan/gonduit"
)

var (
	gonduitURL   = ""
	gonduitToken = ""
	cl           *gonduit.Client
	task         gonduit.Task
)

func TestMain(m *testing.M) {
	gonduitURL = os.Getenv("CONDUIT_URL")
	if gonduitURL == "" {
		fmt.Println("CONDUIT_URL is not defined in environment!")
		os.Exit(1)
	}

	gonduitToken = os.Getenv("CONDUIT_TOKEN")
	if gonduitToken == "" {
		fmt.Println("CONDUIT_TOKEN is not defined in environment!")
		os.Exit(1)
	}

	cl = gonduit.NewClient(gonduitURL, gonduitToken, true)
	task = gonduit.Task{}

	os.Exit(m.Run())
}
