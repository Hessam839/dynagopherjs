package login

import "honnef.co/go/js/dom"

func pass() dom.Node {
	doc := dom.GetWindow().Document()
	password := doc.CreateElement("input")
	password.SetAttribute("type", "password")
	password.SetAttribute("name","password")
	password.SetAttribute("id","password")
	password.SetAttribute("value", "password")
	//password.SetAttribute("style","border-radius:10px; font-size: 16px")

	return password
}
