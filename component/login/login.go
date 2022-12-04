package login

import "honnef.co/go/js/dom"

func Render() dom.Node {
	doc := dom.GetWindow().Document()
	frag := doc.CreateDocumentFragment()
	div := doc.CreateElement("div")
	div.SetAttribute("style","")

	form := doc.CreateElement("form")

	userName := user()

	password := pass()

	form.AppendChild(userName)
	form.AppendChild(password)

	div.AppendChild(form)

	frag.AppendChild(div)
	return frag
}