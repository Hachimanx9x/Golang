package main

import (
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

type nombres struct {
	Nombre string
}

type cuerpo struct {
	Titulo     string
	Id         string
	Contenidos []nombres
}
type Codigo struct {
	Lineas string
}

func index(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("public/index.html"))
	data := cuerpo{
		Titulo: "nombre ",
		Id:     "juego",
		Contenidos: []nombres{
			{Nombre: "JS/Ajax.js"},
			{Nombre: "JS/Teclado.js"},
			{Nombre: "JS/Rectangulos.js"},
			{Nombre: "JS/Mando.js"},
			{Nombre: "JS/bucle1.js"},
			{Nombre: "JS/Dimenciones.js"},
			{Nombre: "JS/inicio.js"},
		},
	}
	t.Execute(w, data)
}

func main() {

	http.HandleFunc("/", index)
	router := mux.NewRouter()

	router.HandleFunc("/", index)
	http.ListenAndServe(":8080", router)
}
