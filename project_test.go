// Copyright 2016 Muhammad Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package gonduit_test

import (
	"github.com/shuLhan/gonduit"
	"testing"
)

func TestProjectCreate(t *testing.T) {
	project := gonduit.Project{
		Fields: gonduit.ProjectFields{
			Name: "Test",
		},
		Members: []string{
			"PHID-USER-bp5wehojb6zk3ivd7xzv",
			"PHID-USER-bqqpi62s6o4dwzmnhpwd",
		},
	}

	e := project.Create(cl)

	if e != nil {
		t.Fatal(e)
	}
}

func TestProjectSearchByName(t *testing.T) {
	_, e := cl.ProjectSearchByName("Proofn Dev/Ops")

	if e != nil {
		t.Fatal(e)
	}
}
