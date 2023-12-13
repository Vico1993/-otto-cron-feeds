package main

import (
	"fmt"
	"os"

	"github.com/Vico1993/Otto-client/otto"
	"github.com/Vico1993/otto-cron-feeds/internal/job"
	"github.com/Vico1993/otto-cron-feeds/internal/service"
	"github.com/Vico1993/otto-cron-feeds/internal/utils"
)

var mainTag = "main"

func main() {
	OttoClient := otto.NewClient(
		nil,
		os.Getenv("OTTO_API_URL"),
	)

	// Notify update if chat present
	if os.Getenv("TELEGRAM_ADMIN_CHAT_ID") != "" {
		version := utils.RetrieveVersion()

		service.NewTelegramService().TelegramPostMessage(
			os.Getenv("TELEGRAM_ADMIN_CHAT_ID"),
			"",
			`🚀 🚀 [CRON-FEEDS] Version: *`+version+`* Succesfully deployed . 🚀 🚀`,
		)
	}

	_, err := job.Scheduler.Every(1).Hour().Tag(mainTag).Do(job.Main, OttoClient)
	if err != nil {
		fmt.Println("Couldn't initiate the main job - " + err.Error())
	}
}
