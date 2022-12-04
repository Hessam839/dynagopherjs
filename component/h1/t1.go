package h1

import (
	"fmt"
	"honnef.co/go/js/dom"
)

func Render(Name string) dom.Node {
	h1 := dom.GetWindow().Document().CreateElement("h1")
	h1.SetInnerHTML(fmt.Sprintf("Hello %s", Name))
	h1.SetAttribute("style","color:red")
	return h1
}