package main

import (
	"fmt"

	"github.com/ToffaKrtek/self-executor/initializers"
)

func init() {
	initializers.LoadEnv()
	initializers.Connect()
	initializers.Run()
}

func main() {
	fmt.Println("Hello World!")
}
