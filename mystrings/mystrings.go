package mystrings

import (
	"fmt"
	"strings"
)
//Strings2 Ingresar todo en mayusculas o minuscular con libreria strings
func Strings2(){
	var text = "Hello world, Hello Platzi, Hello Go"
	fmt.Println(strings.ToUpper(text))
	fmt.Println(strings.ToLower(text))
	// puedes seleccionar cual quieres cambiar con replace pero si pones -1 cambian todos y si pones un entero es el numero fijo
	fmt.Println(strings.Replace(text,"Hello","Hola",-1))
	fmt.Println(strings.Split(text,","))
}
