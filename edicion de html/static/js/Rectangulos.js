function Rectangulo(x, y, ancho, alto) {
	this.x = x;
	this.y = y;
	this.ancho = ancho;
	this.alto = alto;
	this.id =  x+ "r" + y;
	this.insertarDOM();
}

Rectangulo.prototype.insertarDOM = function() {
	var div = '<div id="' + this.id + '"></div>';//<div id ="r0102"></div>
	var html = document.getElementById("juego").innerHTML;//inserta codigo en el div juego
	var color = '#' + Math.floor(Math.random() * 16777215).toString(16);//
	document.getElementById("juego").innerHTML = html + div;// se coge lo del html y se suma otra insercion
	document.getElementById(this.id).style.position = "absolute";
	document.getElementById(this.id).style.left = this.x + "px";
	document.getElementById(this.id).style.top = this.y + "px";
	document.getElementById(this.id).style.width = this.ancho + "px";
	document.getElementById(this.id).style.height = this.alto + "px";
	document.getElementById(this.id).style.backgroundColor = color;//color del rectangulo 
}