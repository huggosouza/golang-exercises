package main

import "fmt"

func main() {
	//var (
	//	idade = 17
	//	nome  = "Hugo Souza"
	//)

	//var altura, idade = 160, 17
	//var nome = "Hugo Souza"

	idade, altura, nome := 17, 160, "Hugo Souza"

	//var idade int = 17
	//var nome string = "Hugo Souza"
	fmt.Print("Name: ", nome, "\nAge: ", idade, "\nHeight: ", altura)
}
