package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"io/ioutil"
	"log"
	"os"
	"path"
)

var (
	configPath string
)

type Config struct {
	ListenAddr string `json:"listen"`
	RemoteAddr string `json:"remote"`
	Password   string `json:"password"`
}

func init() {
	home, _ := homedir.Dir()

	configFilename := ".lightsocks.json"

	if len(os.Args) == 2 {
		configFilename = os.Args[1]
	}

	configPath = path.Join(home, configFilename)
}

func (config *Config) SaveConfig() {
	configJson, _ := json.MarshalIndent(config, "", "		")
	err := ioutil.WriteFile(configPath, configJson, 0644)
	if err != nil {
		fmt.Errorf("save config file %s error: %s", configPath, err)
	}

	log.Printf("save config to file %s success\n", configPath)
}

func (config *Config) ReadConfig() {
	if _, err := os.Stat(configPath); !os.IsNotExist(err) {
		log.Printf("reading config from file %s\n", configPath)
		file, err := os.Open(configPath)
		if err != nil {
			log.Fatalf("open config file %s error: %s", configPath, err)
		}

		defer file.Close()

		err = json.NewDecoder(file).Decode(config)
		if err != nil {
			log.Fatalf("illegal format of JSON config file\n:%s", file)
		}
	}
}
