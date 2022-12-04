package vue

import "github.com/siongui/godom"

func Form(class string, data map[string]interface{}, childes ...*godom.Object) *godom.Object {
	element := godom.Document.CreateElement("form")
	element.Set("class", class)
	element.Set("action", data["action"])

	for _, child := range childes {
		if child != nil {
			element.AppendChild(child)
		}
	}

	return element
}
