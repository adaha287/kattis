// Takes two numbers (N, M) and calculates how many people (totally M people)
// will have to sit in each room (N) if equally distributed
package main

import (
	"fmt"
	"strings"
)

func main() {
	var N, M int
	_, _ = fmt.Scanln(&N)
	_, _ = fmt.Scanln(&M)

	base := M / N
	extra := M % N
	for i := 0; i < extra; i++ {
		fmt.Println(strings.Repeat("*", base+1))
	}
	for i := 0; i < N-extra; i++ {
		fmt.Println(strings.Repeat("*", base))
	}
}
