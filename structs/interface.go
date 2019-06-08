package structs

import "fmt"

type Platzi interface {
	Subscribe(name string)
}

func deferTest(){
	fmt.Println("La funcion InterfaceTest ha culminado")
}

func InterfaceTest()  {
	platziCourse := PlatziCourse{Name: "Go", Slug: "go", Skills: []string{"backend"}}

	platziCareer := new(PlatziCareer)
	platziCareer.Name = "GoCareer"
	platziCareer.Slug = "go"
	platziCareer.Skills = []string{"backend"}

	callSubscribe(platziCourse)
	callSubscribe(platziCareer)

	//DEFER es una funcion o un conjunto de instrucciones que se deben ejecutar al final
	defer deferTest()
}

func callSubscribe(p Platzi)  {
	p.Subscribe("Yohan")
}
