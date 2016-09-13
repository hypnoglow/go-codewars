# go-codewars

[![Build Status](https://travis-ci.org/hypnoglow/go-codewars.svg?branch=master)](https://travis-ci.org/hypnoglow/go-codewars)
[![Go Report Card](https://goreportcard.com/badge/github.com/hypnoglow/go-codewars)](https://goreportcard.com/report/github.com/hypnoglow/go-codewars)
[![GoDoc](https://godoc.org/github.com/hypnoglow/go-codewars?status.svg)](https://godoc.org/github.com/hypnoglow/go-codewars)

A Codewars API client for Go.

## Example

```go
token := "some-api-token"
cw := codewars.NewClient(token)

slug := "printing-array-elements-with-comma-delimiters"
kata, _, err := cw.Katas.GetKata(slug)
if err != nil {
    fmt.Printf("Error: %v\n\n", err)
    } else {
    kataJSON, _ := json.MarshalIndent(kata, "", "  ")
	fmt.Printf("Kata:\n%s\n\n", kataJSON)
}
```

See more complex examples in `examples` folder.


## Documentation

[Codewars API documentation](http://dev.codewars.com/).
