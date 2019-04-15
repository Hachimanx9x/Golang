package main

import (
	"net/http"
	"fmt"
	"html/template"
	"log"

)

	

func login(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(r.Method)
	if r.Method =="GET"{
		t, _ := template.ParseFiles("Archivos html/indes.html")
		t.Execute(w,nil)

	}else{
		fmt.Println(r.Method)
		r.ParseForm()
		fmt.Println("Usuario", r.Form["username"])
		fmt.Println("Password", r.Form["userpassword"])
	}
}
 
func main(){
	//fmt.Fprintf(w, "holi")
	http.HandleFunc("/", login)
err := http.ListenAndServe(":5000", nil)
if err != nil{
	log.Fatal("valimos verga", err)
}}