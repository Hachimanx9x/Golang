package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

type Nombres struct {
	Nombre string
}

type Cuerpo struct {
	Titulo string

	Contenidos []Nombres
}

func index(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("public/index.html"))
	data := Cuerpo{
		Titulo: "juego 2d ",

		Contenidos: []Nombres{
			{Nombre: "js/Ajax.js"},
			{Nombre: "js/Teclado.js"},
			{Nombre: "js/Rectangulos.js"},
			{Nombre: "js/Mando.js"},
			{Nombre: "js/bucle1.js"},
			{Nombre: "js/Dimenciones.js"},
			{Nombre: "js/inicio.js"},
		},
	}
	t.Execute(w, data)
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", index)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	http.ListenAndServe(":8080", router)
}
