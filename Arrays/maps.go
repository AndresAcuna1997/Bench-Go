package main

import "fmt"

func main() {
	//Crear
	urlMap := map[string]string{
		"Google":  "https://google.com",
		"Twitter": "https://twitter.com",
	}

	//Acceder una key en especifico
	fmt.Println(urlMap["Google"])

	//Mutar
	urlMap["LinkedIn"] = "https://likedin.com"

	//Borrar
	delete(urlMap, "Google")
}
