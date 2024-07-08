package main

import (
	"learn/pkg/container"
)

func main() {
	app := container.BuildContainer()
	app.Run()
}
