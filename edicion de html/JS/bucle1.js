//namespace - espacio de nombres
//main loop - bucle principal
//aps - actualizaciones por segundo
//fps - frames por segundo
//callback
//1s = 1000ms

var BuclePrincipal = {
	id_Ejecucion: null,
	ultimoRegistro: 0,
	Aps: 0,
	Fps: 0,
	iterar: function(registro_Temporal) {
		BuclePrincipal.id_Ejecucion = window.requestAnimationFrame(BuclePrincipal.iterar);
        
        BuclePrincipal.Actualizar(registro_Temporal);
        BuclePrincipal.Dibujar();
        
        if(registro_Temporal - BuclePrincipal.ultimoRegistro > 999)
  { BuclePrincipal.ultimoRegistro = registro_Temporal;
    console.log("APS: "+ BuclePrincipal.Aps+ " ll FPS:" + BuclePrincipal.Fps);
   BuclePrincipal.Aps=0;
   BuclePrincipal.Fps=0;
  }
    },
	detener: function() {

	},
	Actualizar: function(registroTemporal) {
        teclado.Reiniciar();
        mando.Actualizar();
   BuclePrincipal.Aps++;
	},
	Dibujar: function(registroTemporal) {
BuclePrincipal.Fps++;
	}
};