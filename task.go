// Copyright 2016 Muhammad Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package gonduit

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

//
// TaskStatus contains the information about task status (e.g. open, close,
// resolved, etc.)
//
type TaskStatus struct {
	Value string `json:"value"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

//
// TaskPriority contains the information about task priority, higher value means
// higher priority.
//
type TaskPriority struct {
	Value       int    `json:"value"`
	Subpriority int    `json:"subpriority"`
	Name        string `json:"name"`
	Color       string `json:"color"`
}

//
// TaskPolicy contains information about which groups can view and edit the
// tasks.
//
type TaskPolicy struct {
	View string `json:"view"`
	Edit string `json:"edit"`
}

//
// TaskFields contains global information of a task.
//
type TaskFields struct {
	Name         string       `json:"name"`
	Description  string       `json:"description"`
	Author       string       `json:"authorPHID"`
	Owner        string       `json:"ownerPHID"`
	Points       string       `json:"points"`
	Space        string       `json:"spacePHID"`
	DateCreate   int          `json:"dateCreated"`
	DateModified int          `json:"dateModified"`
	Status       TaskStatus   `json:"status"`
	Priority     TaskPriority `json:"priority"`
	Policy       TaskPolicy   `json:"policy"`
}

//
// Task contain information for creating, editing or searching Maniphest task.
//
type Task struct {
	ID     int        `json:"id"`
	Type   string     `json:"type"`
	Phid   string     `json:"phid"`
	Fields TaskFields `json:"fields"`

	// Projects contain list of projects that this task belong to.
	Projects []string
}

//
// TaskSearchByName will search the task by their status and name.
// Valid status values are: "all", "open"
//
func (cl *Client) TaskSearchByName(status, name string) (tasks []Task, e error) {
	cl.NewRequest(APITaskSearch)

	cl.request.Add("queryKey", status)
	cl.request.AddConstraint("fulltext", name)

	e = cl.Post()

	if e != nil {
		return nil, e
	}

	if cl.respon.ErrCode != "" {
		return nil, errors.New(cl.respon.ErrCode)
	}

	// Decode the result
	result := SearchResult{}
	e = json.Unmarshal(cl.respon.Result, &result)

	if e != nil {
		return nil, e
	}

	e = json.Unmarshal(result.Data, &tasks)

	if e != nil {
		return nil, e
	}

	for x, task := range tasks {
		fmt.Printf("TASK %d << %+v\n", x, task)
	}

	return tasks, nil
}

//
// setParams will set the parameter required to create task based on information
// on `task` object.
//
func (task *Task) setParams(cl *Client) {
	cl.request.AddParam("title", task.Fields.Name)

	if task.Fields.Description != "" {
		cl.request.AddParam("description", task.Fields.Description)
	}
	if task.Fields.Owner != "" {
		cl.request.AddParam("owner", task.Fields.Owner)
	}
	if task.Fields.Priority.Value != 0 {
		cl.request.AddParam("priority",
			strconv.Itoa(task.Fields.Priority.Value))
	}
	if task.Fields.Status.Value != "" {
		cl.request.AddParam("status", task.Fields.Status.Value)
	}
	if len(task.Projects) > 0 {
		cl.request.AddParams("projects.add", task.Projects)
	}
}

//
// Create will create new task based on information in `task` object.
// User must set,
// - task.Fields.Name for task title
//
// User may set,
// - task.Fields.Description
// - task.Projects for list of project that this belong to
// - task.Fields.Owner
// - task.Fields.Priority.Value
// - task.Fields.Status.Value
//
func (task *Task) Create(cl *Client) (e error) {
	cl.NewRequest(APITaskEdit)

	task.setParams(cl)

	e = cl.Post()

	if e != nil {
		return e
	}

	if cl.respon.ErrCode != "" {
		return errors.New(cl.respon.ErrCode)
	}

	rr, e := cl.respon.DecodeResult()

	task.ID = rr.Object.ID
	task.Phid = rr.Object.Phid

	fmt.Printf("TASK CREATED << %+v\n", task)

	return nil
}
