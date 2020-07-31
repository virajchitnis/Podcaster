package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// Config type for managing config files
type Config struct {
	Server struct {
		Port          string `yaml:"port"`
		Host          string `yaml:"host"`
		WebsiteRoot   string `yaml:"website_root"`
		DataDirectory string `yaml:"data_directory"`
	} `yaml:"server"`
}

func readConfigFile(config string) Config {
	f, err := os.Open(config)
	if err != nil {
		dir := filepath.Dir(config)
		dirInfo, err := os.Lstat(dir)
		if err != nil {
			log.Fatal(err)
		}
		if dirInfo.IsDir() {
			fmt.Println("Creating new configuration file...")
			newConfig := Config{}
			newConfig.Server.Port = "8080"
			newConfig.Server.DataDirectory = "/var/podcaster"

			data, err := yaml.Marshal(&newConfig)
			if err != nil {
				log.Fatal(err)
			}
			werr := ioutil.WriteFile(config, data, 0644)
			if werr != nil {
				log.Fatal(err)
			}
			return newConfig
		}
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		fmt.Println("Invalid config file format!")
		log.Fatal(err)
	}
	return cfg
}
