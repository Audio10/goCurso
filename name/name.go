package name

import "fmt"

// GetName obtiene y retorna el nombre de una persona
func GetName() string{
	var name string
	name = "Sin nombre"
	fmt.Print("Escribe tu nombre a continuacion: ")
	fmt.Scanf("%s\n", &name)
	return name
}
