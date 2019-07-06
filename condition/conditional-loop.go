package condition

import (
	"fmt"
)

// ConditionalLoop is a loop with condition)
func ConditionalLoop() {

	a := 2

	for a < 10 {
		fmt.Println(a)
		a++
	}

	fmt.Println("done;")
}
