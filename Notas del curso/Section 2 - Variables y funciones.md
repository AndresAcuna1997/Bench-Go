
### Variables:

Para declarar variables se puede hacer de dos formas:

```go
var num int = 10 / Forma explicita
num := 10        / Go va a inferir el tipo de la variable num
```

En caso de usar la segunda forma para reasignar la variable se hace sin los dos puntos `:`

Para declarar varias variables en una sola Linea se puede hacer de la siguiente manera:

```go
//Dejando que GO inifera el tipo de la variable
varOne,varTwo, VarThree := 1,2,3

//Forzando el tipo de las variavles
varOne,varTwo, VarThree float64 := 1,2,3
```
## Constantes:

Para declarar un constante se usa la palabra reservada `const`
### Funciones

Para declarar una función se usa la palabra reservada `func` 

```go
func main() {
	//Codigo
}
```

por defecto esta función no retorna algún valor si deseamos retornar algún valor de la función usamos la siguiente sintaxis

```go
func newCard () string {
	return "Ace of Spades"
}

o
//Podemos en la declaracion del return crear las
//variables que van a ser retornadas directamente
func newCard () (num1 int, num2: float64) {
	num1 = 12
	num2 = 21.4
	return num1, num2 
}
```

En caso de que se declaren las variables en la intención de retorno de la function se puede usar solamente la palabra `return` sin especificar los valores después del return.