package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

//some changes here
//some more changes

type Commands struct {
	Cmds []Command `json:"commands"`
}

type Command struct {
	Name string `json:"name"`
	Cmd  string `json:"command"`
}

func main() {
	bot, err := tgbotapi.NewBotAPI(env("TELEGRAM_BOT_TOKEN", ""))
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Connected to bot API")

	commandsFileName := env("COMMANDS_FILE", "commands.json")
	file, err := ioutil.ReadFile(commandsFileName)
	if err != nil {
		log.Panic(err)
	}
	var commands Commands
	err = json.Unmarshal(file, &commands)

	updateConfiguration := tgbotapi.NewUpdate(0)
	updateConfiguration.Timeout = 60

	updates, err := bot.GetUpdatesChan(updateConfiguration)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		if update.Message.IsCommand() {
			for _, command := range commands.Cmds {
				if command.Name == update.Message.Command() {
					out, _ := exec.Command("sh", "-c", command.Cmd).Output()
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, string(out))
					msg.ReplyToMessageID = update.Message.MessageID
					bot.Send(msg)
				}
			}
		}
	}
}

func env(key, defaultValue string) string {
	val := defaultValue
	if envVal := os.Getenv(key); envVal != "" {
		val = envVal
	}
	return val
}
