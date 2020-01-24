go-cli-log
==========

[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/b4b4r07/go-cli-log)

This package provides Terraform-like log system for your CLI tool.
The log level defaults to `TRACE`, `DEBUG`, `INFO`, `WARN`, `ERROR`.

```console
$ go get github.com/b4b4r07/go-cli-log
```

## Example

```go
clilog.Env = "YOUR_APP_LOG"
clilog.SetOutput()

log.Printf("[INFO] run main function")
```

```console
$ YOUR_APP_LOG=trace go run _example/main.go
2020/01/24 20:49:35 [INFO] run main function
```

## License

MIT
