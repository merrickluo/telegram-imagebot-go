package main

import (
	"github.com/messiahluo/tgimagebot/telegramapi"
	"github.com/messiahluo/tgimagebot/searchapi"
)

func main() {
	updatesch := make(chan []telegramapi.Update)

	go telegramapi.StartFetchUpdates(&updatesch)

	for updates := range updatesch {
		for _, update := range updates {
			imageUrl := searchapi.SearchImageForKeyword(update.Message.Text)
			if len(imageUrl) > 0 {
				telegramapi.SendMessage(update.Message.Chat.Id, imageUrl)
			}
		}
	}
}
