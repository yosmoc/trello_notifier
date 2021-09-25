package main

import (
	"log"
	"os"

	"github.com/adlio/trello"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file.")
	}

	client := trello.NewClient(os.Getenv("API_KEY"), os.Getenv("TOKEN"))

	listId := os.Getenv("LIST_ID")
	list, err := client.GetList(listId)

	if err != nil {
		log.Fatalf("error getting list id %v", listId)
	}

	cards, err := list.GetCards()

	if err != nil {
		log.Fatalf("error getting cards")
	}

	for _, card := range cards {
		err := notifier(card.Name, card.URL)
		if err != nil {
			log.Fatal(err)
		}
	}
}
