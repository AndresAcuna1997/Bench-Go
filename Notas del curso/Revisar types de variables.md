En Go podemos revisar el type de una variable y en base a el ejecutar algún código. Esto se puede hacer con un `switch` o con una expresión especial.

#### Switch

```go
switch value.(type) {// value es el nombre de la var puede cambiar
	case int:
	//Codigo si la variable es un Int
	case float64:
	//Codigo si la variable es un float64
	case string:
	//Codigo si la variable es un string
}
```

#### Expresión 

```go
typedVal, ok := value.(int / float / string ...etc)
```

En este caso usar el `.(type)` va a retornar dos valores:
1. `typedVal`: El valor de la variable a la cual se hizo la revision de type
2. `ok`:  un valor `bool` es cual es `true` si el valor de la variable es igual a la puesta dentro de los paréntesis.