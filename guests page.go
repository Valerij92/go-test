package main

import (
	"http"
	"template"
)

const всего = 16

var последняя = 0

type записи struct {
	сообщение [всего]string
}

var гостевая записи

var шаблон *template.Template = template.MustParseFile("шаблон.html", nil)

func увидеть(w http.ResponseWriter, r *http.Request) {
	ошибка := шаблон.Execute(&гостевая, w)
	if ошибка != nil {
		http.Error(w, ошибка.String(), http.StatusInternalServerError)
	}
}

func добавить(w http.ResponseWriter, r *http.Request) {
	if последняя < всего {
		гостевая.сообщение[последняя] = r.FormValue("body")
		последняя++
	}
	http.Redirect(w, r, "/увидеть/", http.StatusFound)
}

func main() {
	http.HandleFunc("/увидеть/", увидеть)
	http.HandleFunc("/добавить/", добавить)
	http.ListenAndServe(":8080", nil)
}
