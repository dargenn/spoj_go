package main

import (
	"fmt"
)

func main() {
	var number int
	for true {
		_, _ = fmt.Scanln(&number)
		if number == 42 {
			break
		}
		fmt.Println(number)
	}
}
