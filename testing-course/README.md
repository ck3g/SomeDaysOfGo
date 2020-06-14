# Golang testing

Hands-on project from Udemy course [Unit, integration and functional Testing in Golang](https://www.udemy.com/share/1020jMA0QaeF1aTHg=/).


## Run tests

```
$ go test ./...
```

### Test coverage

```
$ go test --cover ./...
```

Test coverage takes into account that certain parts of the code have been hit from test code, but it doesn't guarantee that a test result is correct.

### Output debug information from the tests

Use `t.Log` and `t.Logf`. Then run tests with the verbose mode to see the result on the screen.

```
$ go test -v ./...
```
