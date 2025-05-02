/*
Go Lang has two types of packages

  - Executable:
    When compiled, generates an executable file

  - Reusable:
    Packages that are used to support code. A kind of "library" per say

The name of the package is what determines what type it is.
If we use the package 'main', that it means we want an executable file to be created when compiled.

Any other name will create an Reusable Package.
*/
package main //package == project == workspace

/*
fmt (Shortage for "format") is the standard library package for Go

We can find packages at https://golang.org/pkg/

*/
import "fmt"

/*
we always declare functions using the 'func' keyword.
The list of arguments are passed inside the brackets
*/
func main() {
	// since we are creating an executable file by using the 'main' package,
	// we also need to have a 'main' function.
	fmt.Println("hello world!")
}

/*Running code with go:

0: have go installed
1: write code in go lang in a .go file just like this one!
2: once code is finished, go to the terminal and use "go run file.go"
3: code will be run! Profit $$$


important go CLI commands:

	build       compile packages and dependencies
	clean       remove object files and cached files
	fmt         gofmt (reformat) package sources
	get         add dependencies to current module and install them
	install     compile and install packages and dependencies
	run         compile and run Go program
	test        test packages

go build creates an executable file for the current OS (.exe for windows, etc), and does not run it.

*/
