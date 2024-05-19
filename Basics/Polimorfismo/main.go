package main

func main() {
	var p Idioma = Portugues{}
	p.Saudar("Oi")

	var i Idioma = Ingles{}
	i.Saudar("Hello")

	var e Idioma = Espanhol{}
	e.Saudar("Hola")
}
