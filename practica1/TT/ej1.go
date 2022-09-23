package main

import (
	"fmt"
)

func main() {

	name := "Franco"

	fmt.Printf("Tiene %d palabras\n", len(name))

	for i := 0; i < len(name); i++ {
		fmt.Printf("%c\n", name[i])
	}

}
