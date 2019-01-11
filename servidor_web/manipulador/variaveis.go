package manipulador

import "text/template"

//Modelos armazena os modelos de pagina que serao executados pelos manipuladores
var Modelos = template.Must(template.ParseFiles("html/ola.html"))
