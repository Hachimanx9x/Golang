(function(d, axios){ // recibe document y axios
    "use strict";// metodo estricto
    var inputFile = d.querySelector('#inputFile');//busque el campo con la id inputfile
    var divNotification = d.querySelector('#alert');//busque el campo con la id alert

    inputFile.addEventListener('change', addFile);// agregamos el evento cuando cambie el valor se ejecuta la funcion addfile
 
    function addFile(e) {//recive el evento
        var file = e.target.files[0];//es una arreglo donde trae todos los archivos en este caso el primero en file solo se guarda true o falce dado que solo verifica si hay archivo
        if (!file) {//en caso de cancelar por que el archivo no se recibio
            return;//termina la ejecucion de la funcion
        }
        upload(file);//llama la funcion upload
    }

    function upload(file) {//recibe si hay una archivo
        var formData = new FormData();//creamos la variable
        formData.append('file', file);//añadimos('del archivo', archivo recibido)
        
        post('/upload', formData)//url /upload de hay recibe y la variable que tiene la info
            .then(onResponse)//cuando se recibe se ejecuta la funcion onResponce
            .catch(onResponse);//si falla igual se ejecuta
    }

    function onResponse(response) {//recibe la respuesta del server
        var className ='' ;
        if(response.status !== 400)  //si no es igual a 400 osea no hay error
        { className=  'success' //si todo sale bien className es igual a 'success'
     }else{ 
         className=  'error'//sino es igual a error
        }
        
        
        

        divNotification.innerHTML = response.data;//trae el mensaje de server
        divNotification.classList.add(className);//agregamos la clase que es la respuesta
        setTimeout(function () {//funcion sin nombre
            divNotification.classList.remove(className);//removemos el letrero
        }, 3000);//se ejecuta despues de 3 segundos el remover
    }

})(document, axios);//
/*
la sintaxis de una funcion auto ejecutable
(function(){}(); )

El metodo estricto hace 

Elimina algunos errores silenciosos de JavaScript cambiándolos a que lancen errores.
Corrige errores que hacen difícil para los motores de JavaScript realizar optimizaciones: a veces, el código en modo estricto puede correr más rápido que un código idéntico pero no estricto.
Prohibe cierta sintaxis que probablemente sean definidas en futuras versiones de ECMAScript.

*/
