package controllers//nombre de la carpeta

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

// para poder recibir la peticion del cliente y mandar una respuesta
func UploadFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {//si no es motodo post
		http.Redirect(w, r, "/", http.StatusSeeOther)//redirecciona al la ruta por defecto y un codigo de estado
		return
	}

	file, handle, err := r.FormFile("file")//recibe o un archivo un hadle o un error u e formfile recibe el nombre que le diste al
	// formData.append('file', file) la parte 'file' es el nombre el otro el la comprobacion
	if err != nil {//si hay algo en err ocurio un error
		fmt.Fprintf(w, "%v", err.Error())//imprimimos el error
		return//termina
	}
	defer file.Close()//seramos el archivo por que no lo seguimos usandon

	mimeType := handle.Header.Get("Content-Type")// el tipo de archivo que se esta usando Content-Type es el tipo de archivo
	switch mimeType {//mira todos los casos
	case "image/jpeg", "image/jpg", "image/png"://si es igual a alguno de los casos
		saveFile(w, file, handle)//llamamos la funcion mandandole ResponseWriter que es la respuesta del server el archivo y la cabezera 
			default://si no es ninguno
		jsonResponse(w, http.StatusBadRequest, "El formato de la imágen no es válido")//se manda
	}
}
//aqui guardamos el archivo
func saveFile(w http.ResponseWriter, file multipart.File, handle *multipart.FileHeader) {
	//recibe ResponseWriter una variable file usando el paquete multipárt.file y un puntero
	data, err := ioutil.ReadAll(file)//usando el paquete ioutil se lee todos lo archivos del arreglo file
	if err != nil {//si hay algo en error
		fmt.Fprintf(w, "%v", err.Error())//se imprime
		return//
	}

	err = ioutil.WriteFile("./files/"+handle.Filename, data, 0666)//crea el archivo en la ruta .file con el nombre del archivo usando su data en bits usando los permisos del sistema operativo que es 0666
	if err != nil {// si algo salio mal
		fmt.Fprintf(w, "%v", err.Error())//se imprime el error
		return
	}
	jsonResponse(w, http.StatusCreated, "Archivo guardado exitosamente")//si todo salio se manda el letrero que todo salio bien
}
//recibe el ResponseWriter un codigo de estado y un mensaje en string
func jsonResponse(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")//le va a mnadar al cliente de tipo aplicacion json
	w.WriteHeader(code)//escribe en la cabezera el codigo
	fmt.Fprintf(w, message)//imprime en el html es el mensaje usando el ResponseWriter
}