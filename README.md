# gcv-go

[![GoDoc](https://godoc.org/github.com/gagliardetto/gcv-go?status.svg)](https://godoc.org/github.com/gagliardetto/gcv-go)
[![GitHub license](https://img.shields.io/github/license/gagliardetto/gcv-go.svg)](https://github.com/gagliardetto/gcv-go/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/gagliardetto/gcv-go)](https://goreportcard.com/report/github.com/gagliardetto/gcv-go)

## Description

Library written in Golang that provides easy access to the Google Cloud Vision REST API. Google is not affiliated and does not endorse or recommend gcv-go.

## Current Status

Initial version.

## Requirements

+ Go 1.4 or higher
+ Google Cloud account
+ API key

## Installation

```
go get -u github.com/gagliardetto/gcv-go
```

## Getting Started

To get started, create a new project on Google Cloud Console, activate Google Cloud Vision API, and obtain an API key.

To do all of this, you can follow this guide: https://cloud.google.com/vision/docs/getting-started

## Examples

#### Recognize a logo

```go
package main

import (
	"fmt"
	gcv "github.com/gagliardetto/gcv-go"
)

func main() {
	credentials := gcv.Credentials{
		APIkey: "<my api key>",
	}
	client, err := gcv.NewClient(credentials)
	if err != nil {
		fmt.Println(err)
		return
	}

	requestBatch := gcv.RequestBatch{}

	request := gcv.Request{}
	request.AddImageFromFile("google.png")
	request.Features.Add(gcv.LOGO_DETECTION, 1)
	
	requestBatch.Add(request)

	response, err := client.Do(requestBatch)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("This is a %v logo (score: %v)", response[0].LogoAnnotations[0].Description, response[0].LogoAnnotations[0].Score)
}

```