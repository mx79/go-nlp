# go-nlp

## Description

**Go-nlp** is a versatile natural language processing utility.

When I started Golang, I realised that the field of natural language processing was not very developed in this programming language.

So I thought it might be interesting to create a utility package for Gophers.

## Badges

[![forthebadge made-with-go](http://ForTheBadge.com/images/badges/made-with-go.svg)](https://go.dev/)

[![GoDoc reference example](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/mx79/go-nlp)
[![GoReportCard example](https://goreportcard.com/badge/github.com/nanomsg/mangos)](https://goreportcard.com/report/github.com/mx79/go-nlp)

[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/mx79/go-nlp)](https://github.com/mx79/go-nlp)
[![Maintenance](https://img.shields.io/badge/Maintained%3F-yes-green.svg)](https://github.com/mx79/go-nlp/graphs/commit-activity)

## Installation

To import the **go-nlp** package into your project, simply issue the following command: `go get github.com/mx79/go-nlp`

To update the package from time to time, use one of the two commands: `go get -u github.com/mx79/go-nlp`
or `go install github.com/mx79/go-nlp@latest`

## Usage

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
	// My name Max, I will use sentence test. ?!

    // Lemmatizer object
	lemmatizer := clean.NewLemmatizer("en")
	fmt.Println(lemmatizer.Lemm(sentence))
	// My name be Max, I will use this sentence as a test. ?!

	// Stemmer object
	stemmer := clean.NewStemmer("en")
	fmt.Println(stemmer.Stem(sentence))
	// My name is Max, I will use this sentenc as a test. ?!
	

	// Cleaning func
	fmt.Println(clean.Lower(sentence))
	fmt.Println(clean.Tokenize(sentence))
	fmt.Println(clean.RemoveAccent(sentence))
	fmt.Println(clean.RemovePunctuation(sentence))
	// my name is max, i will use this sentence as a test. ?!
	// [My name is Max, I will use this sentence as a test. ?!]
	// My name is Max, I will use this sentence as a test. ?!
	// My name is Max I will use this sentence as a test

	// Purger object that is calling every cleaning package object and func
	p := clean.NewTextPurger("en", true, false, true, true, true, true)
	fmt.Println(p.Purge(sentence))
	// my name max i will use sentenc test
}
```

## Support

This repository is maintained, and you can create a ticket directly on it for any bug or suggestion for improvement.

## Roadmap

In the future, I would like to integrate into this package an intention detection object based on the **RandomForest** algorithm.

Ideally, I would like to integrate as many languages as possible with **stopwords**, **lemmatization** and **stemming** data.

## Contributing

This project is opened to contribution.

## Acknowledgment

I would like to thank ranks.nl for providing me with stopwords for over 30 countries. 
I also thank the contributors to the data in the famous Python NLTK package. 
For this package, I got the stemming data of about ten languages. 