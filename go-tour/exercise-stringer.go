package main

import (
	"fmt"
	"strconv"
)

type IPAddr [4]byte

type I interface {
	M()
}

func (ip *IPAddr) String() string {

	str := fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])

	fmt.Println(str)

	return str
}

func (ip IPAddr) M() {
	fmt.Println(ip)
}

func main() {
	hosts := map[string]*IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)

		fmt.Println(ip.String())
	}

	i, err := strconv.Atoi("42")

	if err != nil {
		fmt.Printf("couldn't convert number: %v\n ", err)
		return
	}

	fmt.Println("converted integer: ", i)
}
