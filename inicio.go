package main
import (
	"html/template"
	"fmt" //esto se usa para imprimir 
 "net/http"   //crear las entradas y definir puerto

)
//contenido
func index( w http.ResponseWriter, r  *http.Request){
	//fmt.Fprintf(w, "holi")
	template, err := template.ParseFiles("Archivos html/indes.html")
	if err != nil{
		fmt.Fprintf(w, "mala suerte perro")
	}else{
	template.Execute(w,nil)	
	}
}

func main(){
	http.HandleFunc("/", index)
	fmt.Println("el server")
	http.ListenAndServe(":5000", nil)
}