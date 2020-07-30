package main

// Config type for managing config files
type Config struct {
	Server struct {
		Port          string `yaml:"port"`
		Host          string `yaml:"host"`
		WebsiteRoot   string `yaml:"website_root"`
		DataDirectory string `yaml:"data_directory"`
	} `yaml:"server"`
}
