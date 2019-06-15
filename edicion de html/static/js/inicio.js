document.addEventListener('DOMContentLoaded', function() {
    Inicio.Iniciar_juego(); //llamamos al juego
	console.log("Javascript funciona");
}, false);

var Inicio={//creamos la variable que tiene el inicio del juego como el llamado del bucle
    Iniciar_juego: function(){
        console.log("Starrr");
        ajax.cargar_archivo("hojas/sin_nombre.json");
     /*   var r = new Rectangulo(0, 0, 100, 100);
		var r2 = new Rectangulo(100, 0, 100, 100);
        var r3 = new Rectangulo(200, 0, 100, 100);*/
        teclado.Iniciar();
        Dimensiones.inicio();
        mando.iniciar();
        Inicio.recargar_tiles();
        BuclePrincipal.iterar();
    },
    
    recargar_tiles: function(){
        document.getElementById("juego").innerHTML=""; 
        for(var y =0; y<Dimensiones.Obtener_tiles_vertical(); y++){
        for(var x =0 ; x< Dimensiones.Obtener_tiles_horizontal(); x++){
            var r= new Rectangulo(x * Dimensiones.Lado_tiles, 
                                  y * Dimensiones.Lado_tiles,            Dimensiones.Lado_tiles, Dimensiones.Lado_tiles); 
        }//fin del for de ancho    
        }//fin del for de altura
    }
};

//DOM = DOCUMENT OBJECT MODEL
//cuando se carga en java pone en la consolo que el java srcip funciono
// ctrl + f5 - recargar limpiando la caché