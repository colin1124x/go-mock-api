# mock api (by golang)

### Install

```
go get github.com/colin1124x/go-mock-api
```

### Usage

```go
// create api map
router := map[string]map[string]string{
    "GET": map[string]string{
        "greet": "hello",
    },
}

// make a quit channel
quit := make(chan bool)

// and run mock server
err := mock_api.Run(":8000", router, quit)

// then do some api request test 

// finally, close mock server
quit <- true

```
