package main

import (
	"fmt"
)

func main() {
	altura := 1.60
	var aberto bool = true
	fmt.Printf("Tipo: %T\nValor: %v\n", altura, altura)
	fmt.Printf("Tipo: %T\nValor: %v", aberto, aberto)
}
