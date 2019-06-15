var teclado = {
	teclas: new Array(),
	Iniciar: function() {
		document.onkeydown = teclado.GuardarTecla;
	},
	GuardarTecla: function(e) {
		teclado.teclas.push(e.key);
		console.log(e.key);
	},
	TeclaPulsada: function(codigoTecla) {
		return (teclado.teclas.indexOf(codigoTecla) !== -1) ? true : false;/*si esta puldada o no indexof es para saber si existe ne el arrray, !== -1 existe ? evalua /si es verdadero sale a true si no false */
	},
	Reiniciar: function() {
		teclado.teclas = new Array();
	}
};