package main

import (
	"fmt"
	template2 "html/template"
	"net/http"
	"os"
	"strings"
	"text/template"
	"time"
)

func main() {
	tmp := newTemp()
	http.HandleFunc("/test", tmp.handle)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

type templates struct {
	template *template.Template
}

func newTemp() *templates {
	dir, _ := os.Getwd()
	_ = dir
	tmp := &templates{}
	tmp.template = template.Must(template.New("html").ParseFiles("./template/index.gohtml"))
	return tmp
}

func (t *templates) handle(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("handler time: %d\n", time.Now().UnixMilli())

	err := t.template.ExecuteTemplate(w, "T", container())
	if err != nil {
		_ = fmt.Errorf("error: %s", err)
	}
}

func container() string {
	fmt.Printf("container time: %d\n", time.Now().UnixMilli())

	div := div(map[string]string{"id": "form"}, "color: red; font-size: 16px",
		input(map[string]string{"placeholder": "username", "name": "username"},
			"font-size: 16px; color: red"),
		input(map[string]string{"placeholder": "username", "name": "username"},
			"font-size: 16px; color: red"),
		input(map[string]string{"placeholder": "username", "name": "username"},
			"font-size: 16px; color: red"))

	return div
}

func div(attributes map[string]string, styles template2.CSS, childs ...string) string {
	fmt.Printf("div time: %d\n", time.Now().UnixMilli())
	src := `<div %s style="%s">%s</div>`
	sb := strings.Builder{}
	for k, v := range attributes {
		sb.WriteString(fmt.Sprintf(`%s="%s" `, k, v))
	}

	return fmt.Sprintf(src, sb.String(), styles, childs)
}

func input(attributes map[string]string, styles template2.CSS) string {
	fmt.Printf("input time: %d\n", time.Now().UnixMilli())
	src := `<input type="text" %s style="%s">`
	sb := strings.Builder{}
	for k, v := range attributes {
		sb.WriteString(fmt.Sprintf(`%s="%s" `, k, v))
	}

	return fmt.Sprintf(src, sb.String(), styles)
}
