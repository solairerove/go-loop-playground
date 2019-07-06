package main

import (
	"fmt"

	nest "github.com/solairerove/go-loop-playground/nest"
)

func main() {

	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}

	nest.NestLoop()
}
