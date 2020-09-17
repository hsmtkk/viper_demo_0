package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("./demo0.yml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	c := Config{}
	if err := viper.Unmarshal(&c); err != nil {
		log.Fatal(err)
	}
	fmt.Println(c)
}

type Config struct {
	SOC struct {
		FeedURL string `mapstructure:"feed_url"`
	} `mapstructure:"soc"`
	WIF struct {
		DirectoryPath string `mapstructure:"directory_path"`
		FileName      string `mapstructure:"file_name"`
	} `mapstructure:"wif"`
}
