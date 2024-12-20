package main

import (
	"fmt"
	"go-to-do/app"
	"go-to-do/configs"
)

func main() {
	configs.Loadconfigs()
	fmt.Println("starting from main")
	app.App()
}
