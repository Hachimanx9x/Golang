package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	rice "github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
)

type Escri struct {
	Texto string
}

func index(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("public/index.html"))
	cambio := Escri{Texto: "js/app.js"}
	w.Header().Set("Content-Type", "application/json") //desimos lo embia es json
	j, err := json.Marshal(cambio)                     //convertimos go a json

	if err != nil {
		panic(err) //en caso de haber error  paras todo XD
	}
	//primero se crea w.header y luego w.writeheader por si algo
	w.WriteHeader(http.StatusOK) //si sobrevivimos a lo anterior se pone que en la cabezera que todo esta bien
	w.Write(j)                   //mandamos el conteniod json
	t.Execute(w, cambio)
}

func main() {
	router := mux.NewRouter()
	router.PathPrefix("/").Handler(http.FileServer(rice.MustFindBox("public").HTTPBox()))
	router.HandleFunc("/", index).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
