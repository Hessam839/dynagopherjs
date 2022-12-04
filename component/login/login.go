package login

import (
	"github.com/siongui/godom"
	"gopherjs/component/vue"
)

func Render() *godom.Object {
	return vue.Form("", map[string]interface{}{"action": "login"},
		vue.Label("", map[string]interface{}{"label": "User Name: "}, nil),
		vue.TextInput("", map[string]interface{}{"name": "username", "id": "username", "placeholder": "username"}, nil),
		vue.Label("", map[string]interface{}{"label": "Password: "}, nil),
		vue.PassInput("", map[string]interface{}{"name": "password", "placeholder": "password"}, nil),
		vue.ButtonInput("", map[string]interface{}{"label": "Login"}, nil),
	)
}
