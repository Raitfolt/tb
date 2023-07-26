package main

import (
	"flag"
	"log"

	tgClient "github.com/Raitfolt/tb/clients/telegram"
	event_consumer "github.com/Raitfolt/tb/consumer/event-consumer"
	"github.com/Raitfolt/tb/events/telegram"
	"github.com/Raitfolt/tb/storage/files"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "storage"
	batchSize   = 100
)

func main() {
	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		files.New(storagePath))

	log.Println("service started")
	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)
	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped")
	}
}

func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for access to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
