"use strict";//usamos el mood estricto

function post(url, data) {//le mandamos una url y la info hacia el server
    return axios.post(url, data)//retorna .then function
        .then(function (response) {// fincion sin nombre le pasosmos el responce
            return response;//lo retornamos
        })
        .catch(function (error) {//en caso de error
            return error.response;//enviamos el mensaje de error
        })
}

/*
Metodo post
Cuando un usuario rellena un formulario en una página web los datos hay que enviarlos de alguna manera. Vamos a considerar las dos formas de envío de datos posibles: usando el método POST o usando el método GET.

*/