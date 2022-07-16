package main

import "fmt"

func main() {
	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, newCard1(), "Six of Spades", Laptop1(), "Queen of Clubs", "Ace of Spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}

func newCard1() string {
	return "King of Hearts"
}
func Laptop1() string {
	return "Predator Helios"
}
