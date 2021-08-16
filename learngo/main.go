package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func main() {
	totalLenght, upperName := lenAndUpper("luca")
	fmt.Println(totalLenght, upperName)
}
