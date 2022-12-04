package vue

import (
	"github.com/siongui/godom"
)

func Label(class string, data map[string]interface{}, child *godom.Object) *godom.Object {
	element := godom.Document.CreateElement("span")
	element.Set("class", class)
	element.AppendChild(godom.Document.CreateTextNode(data["label"].(string)))
	delete(data, "label")

	for k, v := range data {
		element.Set(k, v)
	}

	if child != nil {
		element.AppendChild(child)
	}

	return element
}
