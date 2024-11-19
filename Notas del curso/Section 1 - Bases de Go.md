Go es un lenguaje de **Tipado estático** (lo opuesto a JavaScript)

La estructura de un archivo de Go es la siguiente:

![[Pasted image 20231209115449.png]]
## Packages 

En Go existen dos tipos de packages **Ejecutables** y **Reusables**:

![[Pasted image 20231209110953.png]]

Para poder crear un ejecutable se debe hacer uso del `package main` este es de uso reservado por Go, y debe tener un funcion especial llamada main la cual va a ser ejecutada inmediatamente.

Por otro lado se pueden ver los packages como agrupaciones de archivos de la siguiente manera:

![[Pasted image 20231209114025.png]]

En el caso anterior todos los tres archivos son parte del package main.
## Imports

Permiten traer funcionalidad de otros paquetes a un archivo, estos paquetes pueden ser parte nativa de Go o pueden ser creados por otros desarrolladores. Es un sistema muy parecido a NPM en javascript.

![[Pasted image 20231209113523.png]]

### Ejecución

Existen dos maneras de ejecutar Código en Go `go run --nombre del archivo--` o `go build --nombre del archivo--`

#### Nota: Para poder crear un build en GO es necesario crear un modulo con `go mod init --path--`

1. Run: Compila y ejecuta el archivo
2. Build:  Compila y crea un ejecutable del archivo. Solo se puede crear un ejecutable del package `main`

![[Pasted image 20231209115559.png]]