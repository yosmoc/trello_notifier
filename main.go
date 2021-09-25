package main

import (
	"log"
	"os"

	"github.com/adlio/trello"
	"github.com/joho/godotenv"
)

func main() {
	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load(".env")

		if err != nil {
			log.Fatalf("Error loading .env file.")
		}
	}

	client := trello.NewClient(os.Getenv("TRELLO_API_KEY"), os.Getenv("TRELLO_TOKEN"))

	listId := os.Getenv("TRELLO_LIST_ID")
	list, err := client.GetList(listId)

	if err != nil {
		log.Fatalf("error getting list id %v", listId)
	}

	cards, err := list.GetCards()

	if err != nil {
		log.Fatalf("error getting cards")
	}

	if len(cards) > 0 {
		for _, card := range cards {
			if err := notifier(card.Name, card.URL); err != nil {
				log.Fatal(err)
			}
		}
		return
	}

	if err := notifier("No cards", "Please add a card for list: "+list.Name); err != nil {
		log.Fatal(err)
	}
}
