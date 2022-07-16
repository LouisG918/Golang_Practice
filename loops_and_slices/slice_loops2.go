package main

import "fmt"

func main() {
	cards := []string{"Ace of Spades", newCard(), newNumber(), newNumber2()}
	cards = append(cards, "One of Spades", "Three of Spades", "Four of Spades", "Five of Spades")

	for _, card := range cards {
		fmt.Println(card)
	}

}

func newCard() string {
	return "Two of Spades"
}
func newNumber() string {
	return "1"
}
func newNumber2() string {
	return "2"
}
