var Dimensiones = {
	Ancho: window.innerWidth || document.documentElement.clientWidth || document.body.clientWidth,
	Alto: window.innerHeight || document.documentElement.clientHeight || document.body.clientHeight,
    Lado_tiles:100,
    escala:1,
	inicio: function() {
		window.addEventListener("resize", function(evento) {
			Dimensiones.Ancho = window.innerWidth || document.documentElement.clientWidth || document.body.clientWidth;
			Dimensiones.Alto = window.innerHeight || document.documentElement.clientHeight || document.body.clientHeight;
            Inicio.recargar_tiles();
			console.log("Ancho: " + Dimensiones.Ancho + " | Alto: " + Dimensiones.Alto);
		});
	},
    Obtener_tiles_horizontal: function(){
    var ladoFinal = Dimensiones.Lado_tiles * Dimensiones.escala;
    return Math.ceil((Dimensiones.Ancho - ladoFinal)/ladoFinal);//ceil redondear numero esto esto cirve para saber cuantos rectangulos caben en la pantalla a lo ancho
    
},
     Obtener_tiles_vertical: function(){
    var ladoFinal = Dimensiones.Lado_tiles * Dimensiones.escala;
    return Math.ceil((Dimensiones.Alto - ladoFinal)/ladoFinal);//ceil redondear numero esto esto cirve para saber cuantos rectangulos caben en la pantalla en lo alto
    
},
    Obtener_total_tiles: function(){
        return Dimensiones.Obtener_tiles_horizontal()* Dimensiones.Obtener_tiles_vertical();
    }
};