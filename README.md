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

#### Run face detection

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
	request.AddImageFromFile("<image filepath>")
	request.Features.Add(gcv.FACE_DETECTION, 1)
	
	requestBatch.Add(request)

	response, err := client.Do(requestBatch)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(response[0].FaceAnnotations)
}

```

#### Run landmark detection

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
	request.AddImageFromFile("<image filepath>")
	request.Features.Add(gcv.LANDMARK_DETECTION, 1)
	
	requestBatch.Add(request)

	response, err := client.Do(requestBatch)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(response[0].LandmarkAnnotations)
}

```

#### Run logo detection

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
	request.AddImageFromFile("<image filepath>")
	request.Features.Add(gcv.LOGO_DETECTION, 1)
	
	requestBatch.Add(request)

	response, err := client.Do(requestBatch)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf(response[0].LogoAnnotations)
}

```

#### Run label detection

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
	request.AddImageFromFile("<image filepath>")
	request.Features.Add(gcv.LABEL_DETECTION, 1)
	
	requestBatch.Add(request)

	response, err := client.Do(requestBatch)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(response[0].LabelAnnotations)
}

```

#### Run OCR

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
	request.AddImageFromFile("<image filepath>")
	request.Features.Add(gcv.TEXT_DETECTION, 1)
	
	requestBatch.Add(request)

	response, err := client.Do(requestBatch)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(response[0].TextAnnotations)
}

```

#### Compute image safe-search properties

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
	request.AddImageFromFile("<image filepath>")
	request.Features.Add(gcv.SAFE_SEARCH_DETECTION, 1)
	
	requestBatch.Add(request)

	response, err := client.Do(requestBatch)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(response[0].SafeSearchAnnotation)
}

```

#### Compute a set of properties about the image

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
	request.AddImageFromFile("<image filepath>")
	request.Features.Add(gcv.IMAGE_PROPERTIES, 1)
	
	requestBatch.Add(request)

	response, err := client.Do(requestBatch)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(response[0].ImagePropertiesAnnotation)
}

```