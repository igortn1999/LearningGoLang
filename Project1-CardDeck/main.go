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
	//we could declare the variable like this.
	var card string = "Ace of Spades"
	//However, Go can infere the type based on the value that is being assigned:
	card2 := "Ace of Spades" //Beware: the ':=' only works when CREATING the variable

	//now to reassig...
	card2 = "Five of Diamonds"

	fmt.Print(card, "\n", card2, "\n")

}
