package main

import (
	"fmt"

	condition "github.com/solairerove/go-loop-playground/condition"
	nest "github.com/solairerove/go-loop-playground/nest"
)

func main() {

	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}

	fmt.Println()

	nest.NestLoop()

	fmt.Println()

	condition.ConditionalLoop()
}
