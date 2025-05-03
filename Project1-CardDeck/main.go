package main

import "fmt"

func main() {

	/* declaring a variable

	breaking down:
		var -> means we are declaring a variable
		card -> name of the variable
		string -> type. Make sure only a string will ever be assigned to this variable

	Go is a statically typed language (just like Java and the contrary of JavaScript which is dynamically typed)
	*/

	// we could declare the variable like this.
	var card string = "Ace of Spades"
	// However, Go can infere the type based on the value that is being assigned:
	card2 := "Ace of Spades" //Beware: the ':=' only works when CREATING the variable
	// now to reassign...
	card2 = "Five of Diamonds"

	// using the new function to get a card value
	card3 := newCard()

	/* Array x Slice

	Go Lang has 2 types of arrays:

		Array:
			Fixed length list of things
			Size is static

		Slice:
			Can grow or shrink dependenig on the number of elements it contains
			"Dynamic array"

	For both, all the data inside of it must always be the same of the type.
	*/

	// creating a new Slice of cards
	cards := []string{"Ace of Diamonds", newCard()}
	// now adding a new element to our slice of cards
	cards = append(cards, "7 of Diamonds") //"append" dos not modify the current Slice, it creates a new one

	fmt.Print(card, "\n", card2, "\n", card3, "\n", cards, "\n")

	// now using a for loop to show all the cards in the Slice:
	for i, card4 := range cards { //normal for loop, but we can use "range" to iterate the Slice
		fmt.Print(i, " ", card4, "\n")
	}

	colors := []string{"Red", "Yellow", "Blue"}

	for y, color := range colors {
		fmt.Print(y, " ", color, "\n")
	}

}

// creating a new function to return a new card.
// We need to specify that it returns a string, or it will expect a "void" function
func newCard() string {
	return "Five of Diamonds"
}
