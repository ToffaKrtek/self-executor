package main

import (
	"fmt"

	"github.com/ToffaKrtek/self-executor/initializers"
)

func init() {
	initializers.LoadEnv()
	initializers.Connect()
}

func main() {
	fmt.Println("Hello World!")
}
