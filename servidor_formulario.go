package main

import (
	"net/http"
	"fmt"
	"html/template"
	

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
 
//func main(){

//	http.HandleFunc("/", login)
//err := http.ListenAndServe(":5000", nil)
//if err != nil{
//	log.Fatal("valimos verga", err)
//}}