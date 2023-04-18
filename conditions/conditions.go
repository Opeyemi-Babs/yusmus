package main

import (
	"fmt"
)

func main() {
	/*for i := 0; i < 10; i++ {
		fmt.Printf("Hi\n")
	}
	/* i := 0
	for i < 10 {
		fmt.Printf("Hi\n")
		i++
	} */
	
	/* i := 0
	for i < 10 {
		if i == 5 {break}
		i++

		fmt.Printf("Hi\n")
		i++
	} */
	i := 0
	for i < 10 {
		i++
		if i == 5 {continue}

		fmt.Printf("Hi\n")
		
	}

}
