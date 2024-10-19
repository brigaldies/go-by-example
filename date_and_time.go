package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(fmt.Sprintf("Local tday is %v", time.Now()))
	fmt.Println(fmt.Sprintf("UTC today is %v", time.Now().UTC()))
	fmt.Println(fmt.Sprintf("Truncated UTC today is %v", time.Now().UTC().Truncate(24*time.Hour)))
}
