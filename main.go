package main

import (
	"github.com/siongui/godom"
	"gopherjs/component/login"
)

func main() {
	app := godom.Document.GetElementById("app")
	console := godom.Window.Get("console")
	console.Call("log", "app: ", app)
	app.AppendChild(login.Render())

}
