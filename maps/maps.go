package maps

//GetMap devuelve un map con llaves tipo string y valores tipo entero
func GetMap(name string) int  {
	//CREA UN MAPA O DICCIONARIO CON LLAVES STRING Y VALORES INT

	//mapTest := make(map[string] int)
	mapTest := map[string]int{
		"Juan" : 18,
		"Yohan" : 24,
		"Freddy" : 31,
	}

	mapTest["llave1"] = 1
	mapTest["llave2"] = 100

	//ELIMINA DEL MAP POR LLAVES
	delete(mapTest, "llave1")
	delete(mapTest, "llave2")

	return mapTest[name]
}
