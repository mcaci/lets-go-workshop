package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.Create("./out/text")
	fmt.Fprintln(f, "Hello World!")
	f.Close()
}
