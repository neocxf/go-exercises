package main

import (
	lightsocks "awesomeProject/lightsocks/core"
	"fmt"
)

func main() {
	fmt.Println("we just want to import a package only for invoking it's init method")
	password := lightsocks.RandPassword()

	for _, v := range password {
		fmt.Println(byte(v))
	}
}
