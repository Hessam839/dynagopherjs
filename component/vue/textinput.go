package vue

import "github.com/siongui/godom"

func TextInput(class string, data map[string]interface{}, child *godom.Object) *godom.Object {
	element := godom.Document.CreateElement("input")
	element.Set("type", "text")
	element.Set("class", class)
	for k, v := range data {
		element.Set(k, v)
	}

	if child != nil {
		element.AppendChild(child)
	}

	return element
}
