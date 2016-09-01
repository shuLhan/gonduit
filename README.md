[![GoDoc](https://godoc.org/github.com/shuLhan/gonduit?status.svg)]
(https://godoc.org/github.com/shuLhan/gonduit)
[![Go Report Card](https://goreportcard.com/badge/github.com/shuLhan/gonduit)]
(https://goreportcard.com/report/github.com/shuLhan/gonduit)

`gonduit` is the Go library for working with Phabricator Conduit API [1].

*WARNING*: This library use new (or unstable) API of Conduit.

## Features

* Project
  * Create (`project.edit`): create new project
  * Search (`project.search`): search specific project

* Task
  * Create (`maniphest.edit`): create new task
  * Search (`maniphest.search`): search specific task by status, name, and
    projects
  * IsExist: check wether task is already exist or not

For examples, see the `*_test` files.

## Testing

To run test using your own Phabricator server,
* set the environment `CONDUIT_URL` to your server URL; e.g.
  `http://localhost/`, must with trailing slash, and
* set the environment `CONDUIT_TOKEN`, that you can create at
  `$CONDUIT_URL/settings/user/<username>page/apitokens/`.

[1] https://secure.phabricator.com/conduit/
