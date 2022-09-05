# go-nlp

## Description

Let people know what your project can do specifically. Provide context and add a link to any reference visitors might be
unfamiliar with. A list of Features or a Background subsection can also be added here. If there are alternatives to your
project, this is a good place to list differentiating factors.

**Go-nlp** is a versatile natural language processing utility.

## Badges

[![forthebadge made-with-go](http://ForTheBadge.com/images/badges/made-with-go.svg)](https://go.dev/)

[![GoDoc reference example](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/mx79/go-nlp)
[![GoReportCard example](https://goreportcard.com/badge/github.com/nanomsg/mangos)](https://goreportcard.com/report/github.com/mx79/go-nlp)

[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/mx79/go-nlp)](https://github.com/mx79/go-nlp)
[![Maintenance](https://img.shields.io/badge/Maintained%3F-yes-green.svg)](https://github.com/mx79/go-nlp/graphs/commit-activity)

## Visuals

```go
package main

import (
	"fmt"
	"github.com/mx79/go-nlp/clean"
)

func main() {
	// Test sentence in english
	sentence := "My name is Max and I will use this sentence as a test"

	// Stopwords object
	stopwords := clean.NewStopwords("en")
	fmt.Println(stopwords.Stop(sentence))

	// Lemmatizer object
	lemmatizer := clean.NewLemmatizer("en")
	fmt.Println(lemmatizer.Lemm(sentence))

	// Stemmer object
	stemmer := clean.NewStemmer("en")
	fmt.Println(stemmer.Stem(sentence))

	// Cleaning func
	fmt.Println(clean.Lower(sentence))
	fmt.Println(clean.Tokenize(sentence))
	fmt.Println(clean.RemoveAccent(sentence))
	fmt.Println(clean.RemovePunctuation(sentence))

	// Purger object that is calling every cleaning package object and func
	p := clean.NewTextPurger("en", true, false, true, true, true, true)
	fmt.Println(p.Purge(sentence))
}
```

## Installation

To import the **go-nlp** package into your project, simply issue the following command: `go get github.com/mx79/go-nlp`

To update the package from time to time, use one of the two commands: `go get -u github.com/mx79/go-nlp`
or `go install github.com/mx79/go-nlp@latest`

## Usage

Use examples liberally, and show the expected output if you can. It's helpful to have inline the smallest example of
usage that you can demonstrate, while providing links to more sophisticated examples if they are too long to reasonably
include in the README.

## Support

Tell people where they can go to for help. It can be any combination of an issue tracker, a chat room, an email address,
etc.

## Roadmap

If you have ideas for releases in the future, it is a good idea to list them in the README.

## Contributing

State if you are open to contributions and what your requirements are for accepting them.

For people who want to make changes to your project, it's helpful to have some documentation on how to get started.
Perhaps there is a script that they should run or some environment variables that they need to set. Make these steps
explicit. These instructions could also be useful to your future self.

You can also document commands to lint the code or run tests. These steps help to ensure high code quality and reduce
the likelihood that the changes inadvertently break something. Having instructions for running tests is especially
helpful if it requires external setup, such as starting a Selenium server for testing in a browser.

## Authors and acknowledgment

Show your appreciation to those who have contributed to the project.