// Copyright 2016 Muhammad Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package gonduit

import (
	"fmt"
	"strings"
	"unicode"
)

//
// ConvertToTags will convert list of string to Phabricator tag format, which is
// by replacing non alphanumeric characters with underscores.
//
func ConvertToTags(list []string) (tags []string) {
	sz := len(list)
	if list == nil || sz == 0 {
		return
	}

	tags = make([]string, sz)
	for x, s := range list {
		tmp := make([]rune, len(s))

		for y, r := range strings.ToLower(s) {
			if unicode.IsLetter(r) || unicode.IsNumber(r) ||
				r == '_' {
				tmp[y] = r
			} else {
				tmp[y] = '_'
			}
		}

		tags[x] = string(tmp)
	}

	if DEBUG >= 2 {
		fmt.Printf("[gonduit] ConvertToTags >> %v\n", tags)
	}

	return
}
