package packageA

import "fmt"

type Delegate struct {
}

func (d *Delegate) Execute() {
	// do something

	fmt.Println("delegate executed in packageA...")
}
