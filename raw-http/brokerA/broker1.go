package brokerA

import (
	"github.com/neocxf/go-exercises/raw-http/packageA"
	"github.com/neocxf/go-exercises/raw-http/packageB"
)

type ABroker interface {
	Execute()
}

func New() ABroker {
	return ABroker(new(packageA.Delegate))
}

func New2() ABroker {

	return ABroker(new(packageB.Delegate))
}
