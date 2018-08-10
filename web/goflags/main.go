// build:+linux

// $ go build -x main.go
// $ go build -ldflags="-X main.version=v1.0.0" main.go
// $ go build -ldflags="-w -s" main.go
// $ go tool compile --help
package main

import (
	"fmt"
	"runtime"
)

var version string

func main() {
	fmt.Printf("OS:%v, Architecture:%v\n", runtime.GOOS, runtime.GOARCH)
	fmt.Printf("Version:%v\n", version)
}
