Para mi hay dos reglas importantes:
1. Si dos archivos están dentro del mismo paquete (e.g main) se pueden usar las funciones declaradas dentro de ellas sin restricción.
2. Si deseo usar una función o recurso que es de OTRO paquete se deben seguir los siguientes pasos:
	1. Crear una carpeta para el nuevo paquete
	2. Crear un archivo con los recursos y en el paquete asignarle un nombre (diferente de main)
	3. En el paquete donde deseo importar el otro paquete, importarlo de la siguiente manera
```go
package main
import "example.com/bank/nombre-del-paquete"
```
#### NOTA: Las funciones o recursos expuestos desde el paquete de origen deben estar nombradas con la primera letra en mayúscula

## Paquetes de 3ros

Para usar paquetes de terceros es muy parecido a NPM, su busca un paquete y con la directiva `go get url-del-paquete` lo instalamos. Esto modificara nuestro `.mod` y registrara que paquetes de 3ros estamos usando. 

Si deseamos instalar los paquetes que no poseamos solo usamos la directiva `go get` es algo así como un `npm i`