# autoreadme
forked from github.com/jimmyfrasche/autoreadme


Automatically generate a github README.md for your Go project.

Download:
```shell
go get github.com/1800alex/go-autoreadme
```

If you do not have the go command on your system, you need to [Install Go](http://golang.org/doc/install) first

* * *
Automatically generate a github README.md for your Go project.

autoreadme(1) creates a github-formatted README.md using the same format as godoc(1).
It includes the package summary and generates badges for godoc.org for the complete
documentation and for travis-ci, if there's a .travis.yml file in the directory.
It also includes copy-pastable installation instructions using the go(1) tool.

## HEURISTICS
autoreadme(1) by default imports the Go code in the current directory, unless a directory is specified.

If the -template argument is not given, it tries to use the README.md.template file in the same
directory. If no such file exists, the built in template is used. These rules apply to each
directory visited when -r is specified to run autoreadme(1) recursively. If a README.md already
exists, it fails unless -f is specified.

## EXAMPLES
To create a README.md for the directory a/b/c

```
autoreadme a/b/c
```

To overwrite the README.md in the current directory

```
autoreadme -f
```

To run in the current directory and all subdirectories that contain
Go code

```
autoreadme -r
```

Use the built in template as the basis for a custom template.

```
autoreadme -print-template >README.md.template
```

To override both the default template and a local README.md.template

```
autoreadme -template=path/to/readme.template
```

## TEMPLATE VARIABLES
If you wish to use your own template, These are the fields available to dot:

```
Name - The name of your packages.

Doc - The package-level documentation of your package.

Synopsis - The first sentence of .Doc.

Import - The import path of your package.

RepoPath - The import path without the github.com/ prefix.

Bugs - a []string of all bugs as per godoc.

Library - True if not a command.

Command - True if a command.

Today - The current date in YYYY.MM.DD format.

Travis - True if there is a .travis.yml file in the same directory
	as your package.

Example - a map[name]Example with all examples from the _test files. These
	can be used to include selective examples into the README.
	The Example{} struct has these fields:
		Name - name of the example
		Code - renders example code similar to godoc
		Output - example output, if any
```
