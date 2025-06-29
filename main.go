package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	var data string
	fmt.Scanln(&data)
	fmt.Println("entered: " + data)
}
