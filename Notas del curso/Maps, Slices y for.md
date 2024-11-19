#### Maps: 

Son lo mismo que JS permite tener un struct cuyos keys y values pueden ser cualquier tipo de valor

```go
  //Crear

  urlMap := map[string]string{
    "Google":  "https://google.com",
    "Twitter": "https://twitter.com",
  }

  //Acceder una key en especifico
  fmt.Println(urlMap["Google"])
  
  //Mutar
  urlMap["LinkedIn"] = "https://likedin.com"

  //Borrar
  delete(urlMap, "Google")
```

#### Iterar sobre un Arreglo, slice o map:

Conjunto con la palabra `for` conjunto con la palabra `range`

```go
//Array y slices
for index, value := range mySlice {
 //Your code
}

//Maps
for key, value := range mySlice {
 //Your code
}
```