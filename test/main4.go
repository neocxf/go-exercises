package main

import (
	"path/filepath"
	"fmt"
	"os"
	"log"
	"bufio"
)

func main() {
	pwd, err := filepath.Abs("test/main.go")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	file, err := os.Open(pwd)
	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
}
