package main

import (
	"fmt"
)

func Hello(name string) string {
	return "Hello, " + name
}

func main() {
  hello := Hello("Chris chan")

	fmt.Println(hello)
}
