package main

import "fmt"

func main() {
	card := newCard()
	card1 := newCard1()
	card3 := newCard3()

	fmt.Println(card)
	newCard2()
	fmt.Println(card1)
	fmt.Println(card3)
}

func newCard() string {
	return "Five of Diamonds"
}
func newCard1() string {
	return "Six of Diamonds"
}
func newCard2() {
	deck := 1 + 1
	deck2 := 2 * 2
	deck3 := deck + deck2
	fmt.Println(deck3)
}
func newCard3() int {
	return 5
}
