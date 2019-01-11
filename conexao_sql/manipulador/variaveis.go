package manipulador

import "text/template"

//ModeloOla armazena o modelo de pagina que serao executados pelos manipuladores
var ModeloOla = template.Must(template.ParseFiles("html/ola.html"))

//ModeloLocal armazena o modelo de pagina que serao executados pelos manipuladores
var ModeloLocal = template.Must(template.ParseFiles("html/local.html"))
