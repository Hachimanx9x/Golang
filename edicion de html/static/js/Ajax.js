var ajax = {
  cargar_archivo : function(ruta){
    var peticion = new XMLHttpRequest();
//es un evento
    peticion.onreadystatechange = function(){
/*estados:
0-> Unsent no es iniciada
1-> Opened conectado al server
2-> Header recieved peticion recibida
3-> loding procesando peticion
4-> done peticion finalizada respuesta preparada*/

if(peticion.readyState== XMLHttpRequest.DONE){
if(peticion.status ==200)/*200 que todo esta bien */{
  //recive el json lo vuelve texto a objeto
  console.log(JSON.parse(peticion.responseText));


}else if(peticion.status==400/*error 400*/){
console.log("error")
}else{
  console.log("resultado inesperado o aleatorio")
}
}
    };
    peticion.open("GET",ruta,true);//prepara la peticion
    peticion.send();//envia la peticion
  }
}