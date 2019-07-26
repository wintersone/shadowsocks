package cmd

import (
	"encoding/json"
	"fmt"
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
	home, _ := os.UserHomeDir()

	configFilename := ".shadowsocks.json"

	if len(os.Args) == 2 {
		configFilename = os.Args[1]
	}
	configPath = path.Join(home, configFilename)
}

func (config *Config) SaveConfig() {
	configJson, _ := json.MarshalIndent(config, "", "	")
	err := ioutil.WriteFile(configPath, configJson, 0644)
	if err != nil {
		fmt.Errorf("Saving config to file %s error: %s", configPath, err)
	}
	log.Printf("save config to file %s succeed", configPath)

}

func (config *Config) ReadConfig() {
	if _, err := os.Stat(configPath); !os.IsNotExist(err) {
		log.Printf("loading config from file %s", configPath)
		file, err := os.Open(configPath)
		if err != nil {
			log.Fatalf("cannot open config file: %s", configPath, err)
		}
		defer file.Close()

		err = json.NewDecoder(file).Decode(config)

		if err != nil {
			log.Fatalf("illgal json format:\n %s", file.Name())
		}
	}
}
