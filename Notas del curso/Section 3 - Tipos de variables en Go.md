
### Tipos de variables en Go

Existen 6 tipos de variables nativas en go

1. Boolean  `var a bool = true | a := true`
2. Integer `var a int = 3 | b := 3`
3. float `var a float32 = 3.14 | a := 3.14`
4. string `var a string = "Hola" | a := "Hola"`
5. Array `[_length_]_datatype_{_values_} | [3]int{1,2,3} | [...]int{4,5,6,7,8} `
6. Slice `[]_datatype_{_values_} | []int{1,2,3} | []int{4,5,6,7,8} `

La diferencia entre `Array` y `Slices` es que la longitud de un slice puede variar , un array no.
## Tipos de datos personalizados

Para crear un tipo personalizado se usa la siguiente sintaxis

```go
type deck []string
```

### Receivers en las funciones

Un receivers es por decir una manera de agregar un método a un tipo, algo parecido al prototype en JavaScript, se crea de la siguiente manera

```go

func (d deck) print() {

  for i, card := range d {
    fmt.Println(i+1, card)
  }

}

```


La sintaxis es la de una función, pero antes de el nombre se pone unos paréntesis en los cuales se pasan dos argumentos :

1. El primero hace referencia a la instancia del tipo, se puede pensar como el `this` en un objeto en JavaScript.
2. El segundo hace referencia al tipo sobre el cual se le va agregar el método, en el código anterior, se va a agregar el método `print()` a todas las variables de tipo `deck`

### Slice Range

Para seleccionar en un slice un rango lo podemos hacer por medio de esta sintaxis

```go

var fruits []string{'Apple','Pears','Bananas','Grapes'}

fruits[inicio : Final (sin incluirlo)]

fruits[0:2] // Apple,Pears
fruits[:2] // Apple,Pears, Esta sintaxis es indica desde el inicio del slice hasta el index 2

fruits[2:] // Bananas, Grapes, Desde el index 2 hasta el final del slice 
```

### Retornan y asignar varios valores

Para retornar varios valores podemos usar la siguiente sintaxis

```go

func deal(d deck, handSize int) (deck, deck) {
  return d[:handSize], d[handSize:]
}

```

Después de los argumentos de la función dentro de paréntesis asignamos en orden cuantos valores va a retornar la función y su tipo. Posteriormente para su asignación podemos usar la siguiente sintaxis:

```go
hand, remainingCards := deal(cards, 3)
```

![[Pasted image 20231212162815.png]]

