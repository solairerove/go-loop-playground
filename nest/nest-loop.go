package nest

import (
	"fmt"
)

// NestLoop is a nested loop example
func NestLoop() {

	for i := 0; i <= 10; i++ {
		for y := 0; y < 3; y++ {
			fmt.Println(y)
		}
	}
}
