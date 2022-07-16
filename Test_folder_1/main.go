package main

import "fmt"

func main() {
	card := newCard()
	fmt.Println(card)
	card = "I'm at Kurtis' house"
	fmt.Println(card)
	card1 := OldCard()

	fmt.Println(card1)
	newTable()
	OldTable()
	OlderTable()
	FreshTable()

	card1 = "Louis"
	fmt.Println(card1)
}

func newCard() string {
	return "Five of Diamonds"
}

func newTable() {
	fmt.Println("I want to smoke")
}
func OldTable() {
	fmt.Println("We already smoked")
}
func OlderTable() {
	fmt.Println("We can smoke again")
}
func FreshTable() {
	fmt.Println("We have smoked again")
}
func OldCard() string {
	return "Ace of Spades"
}
