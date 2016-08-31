// Copyright 2016 Muhammad Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package gonduit

import (
	"encoding/json"
	"errors"
	"fmt"
)

//
// ProjectIcon contain icon information used in project.
//
type ProjectIcon struct {
	Key  string `json:"key"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}

//
// ProjectColor contain color information used in project.
//
type ProjectColor struct {
	Key  string `json:"key"`
	Name string `json:"name"`
}

//
// ProjectPolicy contain policy information used in project.
//
type ProjectPolicy struct {
	View string `json:"view"`
	Edit string `json:"edit"`
	Join string `json:"join"`
}

//
// ProjectFields contain global information for project.
//
type ProjectFields struct {
	Name         string        `json:"name"`
	DateCreated  int           `json:"dateCreated"`
	DateModified int           `json:"dateModified"`
	Description  string        `json:"description"`
	Slug         string        `json:"slug"`
	Icon         ProjectIcon   `json:"icon"`
	Color        ProjectColor  `json:"color"`
	Policy       ProjectPolicy `json:"policy"`
}

//
// Project contain information about project in Conduit
//
type Project struct {
	ID      int
	Type    string        `json:"type"`
	Phid    string        `json:"phid"`
	Fields  ProjectFields `json:"fields"`
	Members []string
}

//
// SetName will the the project name to `name`.
//
func (pr *Project) SetName(s string) {
	pr.Fields.Name = s
}

//
// SetPhid will set project PHID value to `s`.
//
func (pr *Project) SetPhid(s string) {
	pr.Phid = s
}

//
// GetName will return project name.
//
func (pr *Project) GetName() string {
	return pr.Fields.Name
}

//
// GetPhid will return project PHID.
//
func (pr *Project) GetPhid() string {
	return pr.Phid
}

//
// ToRequestValues will convert project attribute to URL values.
//
func (pr *Project) ToRequestValues(cl *Client) {
	cl.request.AddParam("name", pr.Fields.Name)
	cl.request.AddParam("description", pr.Fields.Description)

	if pr.Fields.Icon.Key != "" {
		cl.request.AddParam("icon", pr.Fields.Icon.Key)
	}

	if pr.Fields.Color.Key != "" {
		cl.request.AddParam("color", pr.Fields.Color.Key)
	}

	if len(pr.Members) > 0 {
		cl.request.AddParams("members.add", pr.Members)
	}
}

//
// Create new project.
// Required fields: name and description.
//
func (pr *Project) Create(cl *Client) (e error) {
	cl.NewRequest(APIProjectEdit)

	pr.ToRequestValues(cl)

	e = cl.Post()

	if e != nil {
		return e
	}

	if cl.respon.ErrCode != "" {
		return errors.New(cl.respon.ErrInfo)
	}

	rr, e := cl.respon.DecodeResult()

	pr.ID = rr.Object.ID
	pr.Phid = rr.Object.Phid

	if DEBUG >= 1 {
		fmt.Printf("[gonduit] Project.Create << %+v\n", pr)
	}

	return nil
}

//
// ProjectSearchByName will return list of project instance where their name is
// matched with `name`. If no match or error then it will return empty list with
// error value/message.
//
func (cl *Client) ProjectSearchByName(name string) (
	projects []Project,
	e error,
) {
	cl.NewRequest(APIProjectSearch)

	cl.request.AddConstraint("name", name)

	e = cl.Post()
	if e != nil {
		return
	}

	if cl.respon.ErrCode != "" {
		return projects, errors.New(cl.respon.ErrInfo)
	}

	// Decode the result
	result := SearchResult{}
	e = json.Unmarshal(cl.respon.Result, &result)
	if e != nil {
		return
	}

	e = json.Unmarshal(result.Data, &projects)
	if e != nil {
		return
	}

	for x, project := range projects {
		if DEBUG >= 1 {
			fmt.Printf("[gonduit] ProjectSearchByName %d<< %+v\n",
				x, project)
		}
	}

	return
}
