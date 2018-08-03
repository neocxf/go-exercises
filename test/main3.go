package main

import (
	"os"
	"fmt"
	"path/filepath"
)

func main() {
	ex, err := os.Executable()

	if err != nil {
		panic(err)
	}

	exPath := filepath.Dir(ex)

	fmt.Println(exPath)


	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(pwd)

	pwd, err = filepath.Abs("test/main.go")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(pwd)

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	fmt.Println(dir)


	//file, err := os.Open("main.go")
	//
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//fmt.Printf("%+v\n", file)

}
