package main

import (
	"flag"
	. "github.com/keisukefd/TBot/src/Types"
	"github.com/keisukefd/TBot/src/bot"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"strconv"
)

func main() {
	log.Printf("Welcome Bot !")

	var configFile string
	flag.StringVar(&configFile, "config", "config.yml", "configuration file")
	flag.Parse()

	yamlFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatal("Error reading YAML file: ", err)
	}

	var yamlConfig YamlConfig
	err = yaml.Unmarshal(yamlFile, &yamlConfig)
	if err != nil {
		log.Fatal("Error parsing YAML file: ", err)
	}

	var AuthorizedChannels []ChatID
	for _, channel := range yamlConfig.AuthorizedChannels {
		channelInt, _ := strconv.ParseInt(channel, 10, 64)
		AuthorizedChannels = append(AuthorizedChannels, ChatID(channelInt))
	}
	Token := yamlConfig.TelegramToken

	bot.Run(Token, AuthorizedChannels)
}
