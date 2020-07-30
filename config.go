package main

import (
	"fmt"
	"log"
	"os"

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
		fmt.Println("Unable to read config file!")
		log.Fatal(err)
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
