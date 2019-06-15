var mando = {
	objeto: null,//inicia como nada por que no hay nada :p
	eventosDisponibles: 'ongamepadconnected' in window,//existe un evento cuando se conecte un mando devuelve un true o false
	conectado: false,//por defecto
	iniciar: function() {
		if (mando.eventosDisponibles) {
			window.addEventListener("gamepadconnected", mando.Conectar());//
			window.addEventListener("gamepaddisconnected", mando.Desconectar());
		} else {
			mando.Actualizar();
		}
	},
	Conectar: function(evento) {
		mando.objeto = evento.gamepad;//iniciar objeto cuando se conecta el mando
		mando.Identificar();//
	},
	Desconectar: function(evento) {
		console.log("Mando desconectado del índice %d: %s.", evento.gamepad.index, evento.gamepad.id);//cuando se desconecta el mando %d: da el indice %s a la id y el % es para definir varibles inyectadas en la casena
	},
	Actualizar: function() {
		if (!mando.eventosDisponibles) {
			mandos = null;//por si hay mas de uno

			try {
				mandos = navigator.getGamepads ? navigator.getGamepads() : (navigator.webkitGetGamepads ? navigator.webkitGetGamepads : []);//? si no cumple la condicion se me te en el otro ade mas pregunta si esta usando el motor webkit usa gamepad y si hay mas de un control se pone en un array
				mando.objeto = mandos[0];//asigna en el primer indice 
				if(!mando.conectado) {
					mando.conectado = true;
					mando.Identificar();
				}
			} catch(exception) {
				console.log(exception.message);
        	console.log("Control no encontrado");
			}
		}

		if (!mando.objeto) {//si no hay control
			return;//devuelve null y la funcion termina
		}
        if (mando.botonPulsado(mando.objeto.buttons[0])) {//detectar el primer boton
			console.log("Mando: A");
		}

		//continuará
	},
	botonPulsado: function(boton) {
if (typeof(boton) == "object") {//si recibimo un objeto es que llego algo
			return boton.pressed;//devuelve un boollean
		}
		return boton == 1.0;
	},
	Identificar: function() {
console.log("Mando conectado en el índice %d: %s. %d botones, %d ejes.",
			mando.objeto.index, mando.objeto.id, mando.objeto.buttons.length, mando.objeto.axes.length);
	}
};