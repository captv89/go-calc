package main

import (
	_ "embed"
	"fmt"

	"github.com/wailsapp/wails"
)

func Greet(name string) string {
	return fmt.Sprintf("Hello %s!", name)
}

//go:embed frontend/dist/app.js
var js string

//go:embed frontend/dist/app.css
var css string

func main() {

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "calc",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})
	app.Bind(Greet)
	app.Run()
}
