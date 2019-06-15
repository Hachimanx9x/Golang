package main

import ( //paquetes importados
	"encoding/json"
	"log"
	"mux"
	"net/http"
	"strconv"
	"time"
)

//Api rest

type Note struct {
	Title       string    `json:"title"` // alt + 96  es una notacion para paquetes de json
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"create_at"`
}

var noteStore = make(map[string]Note) //acramos un mapa de tipo string guardado en note
var id int

func GetNoteHandler(w http.ResponseWriter, r *http.Request) {
	var notes []Note                  //creamos el objeto que contiene la nota de tipo arrat
	for _, value := range noteStore { //recorremos el map (indice no es usado por eso se queda _)
		notes = append(notes, value) //la nota la vamos agragando al array
	}
	w.Header().Set("Content-Type", "application/json") //desimos lo embia es json
	j, err := json.Marshal(notes)                      //convertimos go a json

	if err != nil {
		panic(err) //en caso de haber error  paras todo XD
	}
	//primero se crea w.header y luego w.writeheader por si algo
	w.WriteHeader(http.StatusOK) //si sobrevivimos a lo anterior se pone que en la cabezera que todo esta bien
	w.Write(j)                   //mandamos el conteniod json

}
func PostNoteHandler(w http.ResponseWriter, r *http.Request) {

	var note Note                                //creamos una variable de tipo Note con su estructura
	err := json.NewDecoder(r.Body).Decode(&note) // descodificar el json que viene en el cuerpo a go
	if err != nil {
		panic(err) //en caso de haber error  paras todo XD
	}
	note.CreatedAt = time.Now()                        //agregamos la fecha a la nota
	id++                                               //aumentamos la id
	k := strconv.Itoa(id)                              //vuelve el numero un string
	noteStore[k] = note                                //notestore en la posicion k se pone los valores de note
	w.Header().Set("Content-Type", "application/json") //desimos lo embia es json
	j, err := json.Marshal(note)                       //convertimos go a json

	if err != nil {
		panic(err) //en caso de haber error  paras todo XD
	}
	//primero se crea w.header y luego w.writeheader por si algo
	w.WriteHeader(http.StatusCreated) //si sobrevivimos a lo anterior se pone que en la cabezera que todo esta bien
	w.Write(j)                        //mandamos el conteniod json
}
func PutNoteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //recibe un puntero a request
	k := vars["id"]     //contiene el id
	var noteUpdate Note
	err := json.NewDecoder(r.Body).Decode(&noteUpdate) //decodifica json a go
	if err != nil {
		panic(err)
	}
	if note, ok := noteStore[k]; ok { // si exite hace la comprobacion
		noteUpdate.CreatedAt = note.CreatedAt //
		delete(noteStore, k)                  //borramos la nota en la posicion k
		noteStore[k] = noteUpdate             //se actualiza la info en la posicion

	} else {
		log.Printf("no encontrado el id %s", k) // %s pone lo que hay en la k
	}
	w.WriteHeader(http.StatusNoContent)

}
func DeleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //recibe un puntero a request
	k := vars["id"]     //contiene el id

	if _, ok := noteStore[k]; ok { // si exite hace la comprobacion

		delete(noteStore, k) //borramos la nota en la posicion k

	} else {
		log.Printf("no encontrado el id %s", k) // %s pone lo que hay en la k
	}
	w.WriteHeader(http.StatusNoContent)
}
func main() {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/api/notes", GetNoteHandler).Methods("GET")
	r.HandleFunc("/api/notes", PostNoteHandler).Methods("POST")
	r.HandleFunc("/api/notes/{id}", PutNoteHandler).Methods("PUT")
	r.HandleFunc("/api/notes/{id}", DeleteNoteHandler).Methods("DELETE")
	server := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("cargando.....")
	log.Fatal(server.ListenAndServe())
}
