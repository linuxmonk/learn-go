## Testing

### Unit Testing

* A file say `http_client.go` would need to have `http_client_test.go` to indicate it to be a test file for `http_client.go`
* The package of `http_client_test.go` can be in the same package as `http_client.go` (say `http_client`) or can be in a separate package like `http_client_test`. 
  - If the test is in the same package then the test can call the internal / unexported functions of the package.
  - If the test is in a separate package (`http_client_test`) then the test can only invoke the exported names. 

### Code Coverage

* 100% on the happy path, about 70-80% on the overall code.
```
$ go test -cover
```

* To generate a cover profile -

```
$ go test -coverprofile c.out
```

* To see the coverage report -

```
$ go tool cover -html c.out
```
