// Copyright 2016 Muhammad Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package gonduit_test

import (
	"fmt"
	"testing"

	"github.com/shuLhan/gonduit"
)

var cl = gonduit.NewClient(CONDUIT_URL, CONDUIT_TOKEN)
var task = gonduit.Task{}

func TestTaskSearchByName(t *testing.T) {
	_, e := cl.TaskSearch("open", "Task", []string{})

	if e != nil {
		t.Fatal(e)
	}
}

func TestTaskSearchByNameProjects(t *testing.T) {
	_, e := cl.TaskSearch("all",
		"Cleaning up all machines and update all required packages",
		[]string{
			"Proofn Dev/Ops",
		})

	if e != nil {
		t.Fatal(e)
	}
}

func TestTaskIsExist(t *testing.T) {
	exist := cl.TaskIsExist("all",
		"Cleaning up all machines and update all required packages",
		[]string{
			"Proofn Dev/Ops",
		})

	if exist {
		fmt.Println("Task is exist")
	} else {
		fmt.Println("Task is not exist")
	}
}

func TestTaskCreate(t *testing.T) {
	task.SetName("Task 1")

	e := task.Create(cl)

	if e != nil {
		t.Fatal(e)
	}
}

func TestTaskCreateWithProjects(t *testing.T) {
	task.SetName("Task 1")
	task.AddProject("Test")

	e := task.Create(cl)

	if e != nil {
		t.Fatal(e)
	}
}

func TestTaskCreateWithPriority(t *testing.T) {
	task.SetName("Task 1")
	task.SetPriority(80)

	e := task.Create(cl)

	if e != nil {
		t.Fatal(e)
	}
}

func TestTaskCreateWithStatus(t *testing.T) {
	task.SetName("Task 1")
	task.SetStatus("invalid")

	e := task.Create(cl)

	if e != nil {
		t.Fatal(e)
	}
}

func TestTaskCreateWithComments(t *testing.T) {
	task.SetName("Task 1")
	task.AddComment("FIRST!!")
	task.AddComment("SECOND!!")

	e := task.Create(cl)

	if e != nil {
		t.Fatal(e)
	}
}
