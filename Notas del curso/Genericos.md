En una manera de decirle a GO que no sabemos que type va a llegar a la función y / o regresar la misma hasta que esta sea llamada

```go
func add[T any](a, b T) T {
	return a+b
}
```

En el caso anterior le decimos que el Type `T` es de `any` ósea cualquier tipo. Esto es problemático ya que no podemos suma dos structs por ejemplo. Para limitarlo a ciertos tipos se puede hacer de la siguiente manera

```go
func add[T int | float64 | string](a, b T) T {
	return a+b
}
```

Ahora Go sabe que `T` puede ser int, float y string los cuales permiten esta operación