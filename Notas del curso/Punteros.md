Variables que guardan una direction en lugar de un valor. Con esto podemos evitar copias innecesarias de valores y podemos mutar directamente un valor.

Para crear un puntero se hace usando el carácter `&` después de la variable a la cual se va a apuntar la dirección en memoria

```go
var age int = 27
var agePointer = &age
```

si se usa agePointer el resultado va a ser la dirección en memoria algo como `0xc00000a0b8` , si se desea usar el valor de la variable en esa dirección se debe "parsear" esto se logra usando el carácter `*` antes del puntero.

```go
var age int = 27
var agePointer = &age

fmt.Print(agePointer) //0xc00000a0b8
fmt.Primt(*agePointer) //27
```

### Editar valores en memoria:

Con punteros podemos editar una valor directamente en memoria esto se hace "de referenciando" el puntero, modificando el valor y luego nuevamente asignándolo al puntero dereferenced

```go
func main() {
  age := 27
  agePointer := &age
  editAgeToAdultYears(agePointer)
  fmt.Print(age) //9
}

func editAgeToAdultYears(age *int){
	*age = *age -18
}
```

El código anterior va a modificar el valor en memoria lo cual significa que el valor de `age` después de llamar la function y pasar el puntero va a ser de 9