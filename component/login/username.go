package login

import "honnef.co/go/js/dom"

func user() dom.Node {
	doc := dom.GetWindow().Document()
	userName := doc.CreateElement("input")
	userName.SetAttribute("type", "text")
	userName.SetAttribute("name","username")
	userName.SetAttribute("value", "username")
	//userName.SetAttribute("style","border-radius:10px; font-size: 16px")

	return userName
}
