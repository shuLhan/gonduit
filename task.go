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

	// Comments contain list of comment in task.
	Comments []string
}

//
// SetName will set task name (or title) to `s`
//
func (t *Task) SetName(s string) {
	t.Fields.Name = s
}

//
// SetDescription will set task description to `s`
//
func (t *Task) SetDescription(s string) {
	t.Fields.Description = s
}

//
// SetOwner will set task owner to `s`.
//
func (t *Task) SetOwner(s string) {
	t.Fields.Owner = s
}

//
// SetStatus will set task status to `s`.
//
func (t *Task) SetStatus(s string) {
	t.Fields.Status.Value = s
}

//
// SetPriority will set task priority to `s`.
//
func (t *Task) SetPriority(v int) {
	t.Fields.Priority.Value = v
}

//
// AddProject will assign project `s` to task only if `s` is not already in
// list.
//
func (t *Task) AddProject(s string) {
	for _, p := range t.Projects {
		if p == s {
			return
		}
	}
	t.Projects = append(t.Projects, s)
}

//
// AddComment will add new comment to task only if `s` is not already in
// list.
//
func (t *Task) AddComment(s string) {
	for _, p := range t.Comments {
		if p == s {
			return
		}
	}
	t.Comments = append(t.Comments, s)
}

//
// GetName will return task name.
//
func (t *Task) GetName() string {
	return t.Fields.Name
}

//
// TaskSearch will search the task by their status, name, and projects.
// Valid status values are: "all", "open".
//
func (cl *Client) TaskSearch(status, name string, projects []string) (
	tasks []Task,
	e error,
) {
	cl.NewRequest(APITaskSearch)

	cl.request.Add("queryKey", status)
	cl.request.AddConstraint("fulltext", name)

	if len(projects) > 0 {
		projects = ConvertToTags(projects)

		cl.request.AddConstraints("projects", projects)
	}

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
		if DEBUG >= 1 {
			fmt.Printf("[gonduit] TaskSearch %d << %+v\n", x,
				task)
		}
	}

	return tasks, nil
}

//
// TaskIsExist will return true if after searching task with `status`, `name`,
// and list of `projects` it return non nil or empty tasks; otherwise it return
// false.
//
func (cl *Client) TaskIsExist(status, name string, projects []string) bool {
	tasks, e := cl.TaskSearch(status, name, projects)
	if e != nil || tasks == nil || len(tasks) == 0 {
		return false
	}
	return true
}

//
// setParams will set the parameter required to create task based on information
// on `task` object.
//
func (t *Task) setParams(cl *Client) {
	cl.request.AddParam("title", t.Fields.Name)
	cl.request.AddParam("priority",
		strconv.Itoa(t.Fields.Priority.Value))

	if t.Fields.Description != "" {
		cl.request.AddParam("description", t.Fields.Description)
	}
	if t.Fields.Owner != "" {
		cl.request.AddParam("owner", t.Fields.Owner)
	}
	if t.Fields.Status.Value != "" {
		cl.request.AddParam("status", t.Fields.Status.Value)
	}
	if len(t.Projects) > 0 {
		cl.request.AddParams("projects.add", t.Projects)
	}
	if len(t.Comments) > 0 {
		for _, c := range t.Comments {
			cl.request.AddParam("comment", c)
		}
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
func (t *Task) Create(cl *Client) (e error) {
	cl.NewRequest(APITaskEdit)

	t.setParams(cl)

	e = cl.Post()

	if e != nil {
		return e
	}

	if cl.respon.ErrCode != "" {
		return errors.New(cl.respon.ErrInfo)
	}

	rr, e := cl.respon.DecodeResult()

	t.ID = rr.Object.ID
	t.Phid = rr.Object.Phid

	if DEBUG >= 1 {
		fmt.Printf("[gonduit] Task.Create << %+v\n", t)
	}

	return nil
}
