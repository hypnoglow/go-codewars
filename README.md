# go-codewars

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
