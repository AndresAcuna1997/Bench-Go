Los structs son básicamente objetos de JS se usan para agrupar información que esta relacionada de alguna manera.

#### Como definir un Struct

Para crear un struct primero debemos crear un `type` personalizado y a este asignarle el tipo struct

```go
type User struct {}
```

Esto se hace fuera de la función main y en caso de que el Tipo vaya ser usado en otro modulo debe ser nombrado Capital case.

Después de esto definimos los campos que va a tener este tipo de struct muy parecido a una interfaz de TS

```go
type User struct {
	nombre string
	appellido string
	edad int
	fechaNacimiento time.Time
}
```

#### Como crear un struct

Para crear un struct se crea una variable y se usa el nombre del `type` junto con las propiedades o vacío

```go
func main() {

	nombre := "Andres"
	appellido := "Acuna"
	edad := 27
	fechaNacimiento Date.now()

	var appUser User = User{} // Struct vacio
	var appUser User = User //Struct con valores
	{nombre, apellido, edad, fechaNacimiento,}
}
```

##### Diferentes maneras de llenar un struct

Se puede llenar un struct de 2 maneras declarando la llave siendo mas explicito o asegurandose de que el orden sea el correcto

1. Explicito: Se declara la key y el value
```go
var appUser User = User 
	{nombre:nombre, apellido:apellido, edad:edad, fechaNacimiento:fechaNacimiento,}
```
2. Orden: Si tenemos todos (?) lo valores del struct podemos simplemente usar los valores si declaran la key, pero debemos asegurarnos de que están en orden
```go
var appUser = User{nombre,apellido,edad,fechaNacimiento}
```

#### Métodos en Structs

Similar a JS una función dentro de un objeto se considera un método, esto se vincula usando un `reciver` que son unos `()` depues del `func` y antes del nombre de la function, en el se denota el nombre para acceder a las propiedades del struct y el `type` al cual va estar ligado.

```go
type User struct {
  firstName string
  lastName  string
  birthdate string
  createdAt time.Time
}

func (u User) outPutUserDetails() { //Se va a acceder con la letra u y esto es un metodo de el type User
  fmt.Println(u.lastName, u.firstName, u.birthdate)
}
```

para ejecutarlo simplemente se llama como una propiedad mas de la instancia de struct 

```go
appUser.outPutUserDetails()
```

#### Mutation methods

Para métodos que van a mutar el struct original es necesario para un puntero para que modifique el valor en memoria ya que en Go siempre se pasa una copia del argumento

```go
func (u *User) clearName() {
  u.firstName = ""
  u.lastName = ""
}
```

Con esto nos aseguramos de mutar los datos originales.
#### NOTA: Para usar un puntero en un Struct no es necesario des referenciarlo para acceder a la data, es un shortcut aceptado por GO

#### Función constructora:

La Función constructora no es una feature de GO, es simplemente una función que centraliza la creación de structs, con esto podemos hacer validaciones en un solo lugar.

```go
func NewUser(firstName, lastName, birthdate string) (*User, error) {

  if firstName == "" || lastName == "" || birthdate == "" {
    return nil, errors.New("empty arguments are invalid")
  }

  return &User{
    FirstName: firstName,
    LastName:  lastName,
    Birthdate: birthdate,
    createdAt: time.Now(),
  }, nil
}
```
#### Exportar Struct y types:

Para exportar structs es igual que los demás elementos de GO, simplemente se usa la primera letra en mayúsculas. Con los tipos es igual pero si no se usa una función constructora las el Nombre del tipo y las propiedades su primera letra debe ir en mayúscula. Esto sirve para limitar que propiedades son expuestas a otros packages.

```go
type User struct {
  FirstName string
  LastName  string
  Birthdate string
  createdAt time.Time
}
```

En el caso anterior todas las propiedades excepto de createdAt son expuestas a otros packages

#### Struct Anidados:

Al igual que TS un tipo o interfaz puede extender las propiedades y métodos de otra interfaz o tipo. En GO se puede hacer lo mismo podemos crear un type de struct en base a otro.

```go
type User struct {
  FirstName string
  LastName  string
  Birthdate string
  createdAt time.Time
}

type Admin struct {
  email string
  password  string
  ===============
  User
  user: User
}
```

Para que el nuevo tipo herede del padre debemos referenciarlo en el type del hijo, se puede referenciar de 2 maneras:
1. Anonima: Solo se pone el type del otro Struct (solo `User` en este caso)
2. Nombrado: Se le agrega al igual que las otras propiedades un nombre (`user User`)
La diferencia entre los dos en como el type hijo va acceder a los métodos del padre
```go
  admin.OutPutUserDetails()  //Anonimo
  admin.user.OutPutUserDetails()  //Nombrado
```
###### Crear Struct anidado:

Se puede usar el mismo patron de una función creadora con una pequeño cambio. En el caso de `Admin` las propiedades nuevas son `email, password` estas se declaran igual. Las propiedades heredadas del padres se declaran de la siguiente manera 

```go
  return Admin{
    email:    email,
    password: password,
    User: User{
      FirstName: "ADMIN",
      LastName:  "GOD",
      Birthdate: "00/00/0000",
      createdAt: time.Now(),
    },
  }
```

Básicamente lo que se hace es crear un struct dentro de un struct, en el ejemplo se esta estableciendo la propiedad `User` un struct de type `User`