package main

import (
	"fmt"
	"log"
	"net/http"

	"controllers"
)

func main() { //se crea en la direccion / se sirva del directorio public
	http.Handle("/", http.FileServer(http.Dir("./public")))
	//usa la funcion de uploadfile
	http.HandleFunc("/upload", controllers.UploadFile)
	//imprimimos que el server funciona
	fmt.Println("Server running")
	//se habilita el puerto 8080 y el log.fatal es por si hay un error que lo imprima
	log.Fatal(http.ListenAndServe(":8080", nil))
}
