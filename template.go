//Copyright 2013 James Frasche. All rights reserved.
//Use of this code is governed by a BSD-License found in the LICENSE file.

package main

import "text/template"

var tmplraw = `# {{.Name}}{{/*
*/}}{{if .Library}} [![GoDoc](https://godoc.org/{{.Import}}?status.svg)](https://godoc.org/{{.Import}}){{end}}{{/*
*/}}{{if .Travis}} [![Build Status](https://travis-ci.com/{{.RepoPath}}.png?branch=master)](https://travis-ci.com/{{.RepoPath}}){{end}}
{{.Synopsis}}

Download:
` + "```" + `shell
go get {{.Import}}
` + "```" + `
{{if .Command}}
If you do not have the go command on your system, you need to [Install Go](http://golang.org/doc/install) first
{{end}}
* * *
{{.Doc}}
{{if .Bugs}}
# Bugs
{{range .Bugs}}* {{.}}{{end}}
{{end}}

{{if .Example}}
# Examples
{{range .Example}}
{{.Name}}
{{.Code}}
{{end}}
{{end}}
`
var tmpl = template.Must(template.New("").Parse(tmplraw))
