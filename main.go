package main

import (
	"alias-method/method"
	"fmt"
)

/**
* Cost:
* - Preprocssing O(n)
* - Generate probability O(1)
 */
func main() {
	input := []float32{0.5, 0.5}

	method.New(input)
	fmt.Println(method.Generate())
	fmt.Println(method.Generate())
	fmt.Println(method.Generate())
}
