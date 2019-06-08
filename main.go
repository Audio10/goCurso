package main

import (
	"fmt"
	"github.com/Audio10/gocurso/numbers"
)

const helloWord string = "Hola %s %s, bienvendio al facinante mundo de go.\n"
const testConst = "Test"

func main() {
	
	/*lastname := "<modificar con mi apellido>"
	firstName := name.GetName()
	number := numbers.Sum(50,50)
	a, b, c := numbers.GetVariables()
	f32, f64 := numbers.GetFloat()
	stringUTF8 := name.GetUnicode()
	fmt.Printf(helloWord, firstName,lastname)
	fmt.Print("Hola mundo")
	fmt.Println(number,a,b,c)
	fmt.Println(f32,f64)
	fmt.Println("Cadena con UFT8" , stringUTF8)
	// VALOR ASCCI Y CON string() aplica el cast a string
	fmt.Println( string("hola"[0]) )
	fmt.Println("Cantidad de letras que tiene hola ", len("hola"))
	structs.GetArray()
	structs.GetSlice()
	flow.IfTest()
	ciclos.ForTest()
	mystrings.Strings2()
	flow.SwitchTest()
	fmt.Println(maps.GetMap("Yohan"))

	CREACION DE UNA STRUCTURA TIPO PLATZICOURSE
	platziCourse1 := new(PlatziCourse)
	platziCourse1.Name = "Go1"
	platziCourse1.Slug = "go1"
	platziCourse1.Skills = []string{"backend"}*/

	//structs.InterfaceTest()

	number, err := numbers.Sum(50, 50)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(number)
}



