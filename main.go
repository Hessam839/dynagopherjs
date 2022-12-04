package main

import (
	"github.com/gopherjs/gopherjs/js"
	"gopherjs/component/h1"
	"gopherjs/component/login"
	"honnef.co/go/js/dom"
)

func print(args ...interface{}) {
	js.Global.Get("document").Call("write",args)
}

func main() {
	app := dom.GetWindow().Document().GetElementByID("app")
	app.AppendChild(h1.Render("Hessam Hashemi"))
	app.AppendChild(login.Render())
	dom.GetWindow().Console().Call("log", "app variable",app)
	print("Hello", "World!")
}