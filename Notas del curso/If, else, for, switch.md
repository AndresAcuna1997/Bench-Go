
### IF

En go los if se declaran de manera similar a JavaScript y poseen las misma validaciones

```go
if num == 2 {
 fmt.Print("Do something")
}
```

A diferencia de JavaScript no existe validación por valor y tipo `==` y `===` solo existe el doble igual `==`

al igual que JavaScript al final de este se puede poner un `else if` para evaluar una segunda / tercera condición si la anterior no fue cumplida.

### For

A diferencia de otro lenguajes de programación para hacer bucles Go solo posee la directiva `for`, para compensar esta es fexible.

1. For estándar: Este for crea una una variable, va a revisar la condición para ejecutar el código dentro del for, y va a aumentar la variable
```go
for i:=0; i < 10; i++ {
//Codigo en bucle
}
```

2. For infinito: Para indicar un for infinito solo se usa la palabra `for` y el código dentro de las llaves.
3. For Condicional: Junto con la palabra `for` se puede poner una condición que retorne un booleano. En este caso el loop seguirá ejecutándose hasta que la condición adyacente al for sea `false`
```go
var num := 1
for num<=10 {
	num +=1
}
```

### Switch

Es igual a Javascript, con la exception que solo un caso se puede ejecutar a la vez por ende el uso de la palabra `break` es redundante. Si el switch esta dentro de un for la palabra reservada `break` ya no servirá para romper el `for` si no para salir del `switch`.

```go
choice := 3
switch choice {
	case 1:
	//Do Something
	case 2:
	//Do Something
	case 3:
	fmt.Print("Case 3")
	default:
	//Do Something
}

```