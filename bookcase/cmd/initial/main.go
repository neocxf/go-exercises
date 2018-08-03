package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/gobuffalo/packr"
	"bytes"
)

func init() {
	box := packr.NewBox("../../public")

	config := box.Bytes("config.json")

	viper.SetConfigType("json")

	err := viper.ReadConfig(bytes.NewBuffer(config))

	if err != nil {

		panic(err)
	}

	if viper.GetBool(`debug`) {
		fmt.Println("service run in debug mode")

	}

	for _, elem := range viper.AllKeys() {
		fmt.Printf("%v : %v\n", elem, viper.GetString(elem))
	}

	fmt.Println(viper.GetString("database.pass"))
}

func main() {

	fmt.Println("hello bookcase initial")
}
