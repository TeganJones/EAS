package main

import (
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))

	http.HandleFunc("/", home)

	//http.HandleFunc("/admin", admin)
	http.HandleFunc("/mail", mailHandler)

	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public/"))))

}

func renderTemplate(res http.ResponseWriter, name string, data interface{}) {
	err := tpl.ExecuteTemplate(res, name, data)
	if err != nil {
		http.Error(res, err.Error(), 500)
	}
}
