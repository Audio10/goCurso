package ciclos

import "fmt"

//ForTest utiliza ciclos for para imprimir del 0 al 5, decrementar de 10 en 10 y terminar un bucle con 1000 iteraciones.
func ForTest() {
	for i := 0; i < 5; i++ {
		fmt.Println("[FOR] ", i)
	}

	c := 100
	for c > 0 {
		c -= 10
		fmt.Println("[FOR] Solo con una condición de c > 0, ", c)
	}

	s := 1000
	for {
		s -= 1
		if s == 0 {
			fmt.Println("Termina el for 'infinito'")
			break
		}
	}
}

