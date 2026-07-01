package main

import (
	"fmt"
	"main/internal/config"
)

func init() {
	config.Setup()
}

func main() {
	fmt.Printf("%+v\n", config.Setting)
}
