// Copyright 2016 Muhammad Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package gonduit_test

import (
	"fmt"
	"testing"

	"github.com/shuLhan/gonduit"
)

func TestConvertToTags(t *testing.T) {
	list := []string{
		"Project Dev/Ops",
		"Project-Mobile-iOS",
	}

	tags := gonduit.ConvertToTags(list)

	fmt.Printf("%v\n", tags)
}
