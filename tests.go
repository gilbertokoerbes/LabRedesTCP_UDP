package main

import (
	"fmt"
	"os"
)

func main() {
	fileopen, _ := os.ReadFile("./file.txt")
	fmt.Print(string(fileopen))
}
