package main

import "fmt"

func hello(name string) string{
	message := fmt.Sprintf("hi %v. welcome!\n", name)

	return message
}

func main(){
	fmt.Printf(hello("opeyemi"))
}