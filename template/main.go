package main

import (
	"fmt"
	"github.com/aymerick/raymond"
)

func main() {
	src := `<input type="text" placeholder="{{ph}}" name="{{name}}">`
	tpl, _ := raymond.Parse(src)
	res, err := tpl.Exec(map[string]interface{}{"ph": "hello", "name": "username"})
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
