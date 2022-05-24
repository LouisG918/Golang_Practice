package main

import "fmt"

// func main() {
// 	//line number 7 blue word means variable
// 	//anything inside of "qoutaiton marks" is called a STRING
// 	card := "Ace of Spades"
// 	card = "Five of Diamonds"
// 	balloon := "Red One"
// 	animal := "Giraffe"
// 	weapon := "Longsword"
// 	Sum := 1 + 1

// 	fmt.Println(card)
// 	fmt.Println(balloon)
// 	fmt.Println(animal)
// 	fmt.Println(weapon)
// 	fmt.Println(Sum)
// }

func main() {
	cards := []string{"Ace of Diamonds", "King of Hearts", "Ace of Spades", newCard()}

	fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds"
}
