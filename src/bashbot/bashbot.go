package main

import (
    "os"
    "os/exec"
    "log"
    "github.com/go-telegram-bot-api/telegram-bot-api"
)


func main() {
    bot, err := tgbotapi.NewBotAPI(env("TELEGRAM_BOT_TOKEN", ""))
    if err != nil {
        log.Panic(err)
    }
    log.Printf("Connected to bot API")

    updateConfiguration := tgbotapi.NewUpdate(0)
    updateConfiguration.Timeout = 60

    updates, err := bot.GetUpdatesChan(updateConfiguration)
	for update := range updates {
		if update.Message == nil {
			continue
		}

        chatName := strings.TrimSpace(update.Message.Chat.Title)
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
	}
}

func env(key, defaultValue string) string {
    val := defaultValue
    if envVal := os.Getenv(key); envVal != "" {
       val=envVal
    }
    return val
}
