// Copyright 2016 Muhammad Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package gonduit_test

import (
	"github.com/shuLhan/gonduit"
	"testing"
)

var cl = gonduit.NewClient(CONDUIT_URL, CONDUIT_TOKEN)

func TestTaskSearchByName(t *testing.T) {
	_, e := cl.TaskSearchByName("open", "Task")

	if e != nil {
		t.Fatal(e)
	}
}

func TestTaskCreate(t *testing.T) {
	task := gonduit.Task{
		Fields: gonduit.TaskFields{
			Name: "Task 1",
		},
	}

	e := task.Create(cl)

	if e != nil {
		t.Fatal(e)
	}
}

func TestTaskCreateWithProjects(t *testing.T) {
	task := gonduit.Task{
		Fields: gonduit.TaskFields{
			Name: "Task 1",
		},
		Projects: []string{
			"Test",
		},
	}

	e := task.Create(cl)

	if e != nil {
		t.Fatal(e)
	}
}
