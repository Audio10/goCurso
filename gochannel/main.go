package main
import "fmt"

func main(){
	ch := make(chan string)
	go func() {
		defer close(ch) // CERRANDO EL CANAL CON close(ch) cuando esta funcion se ejecuta ya no puedes asignar mas valores al canal
		ch<-"Hola Channel"
	}() // este parentecis asigna que cuando se acabe la funcion se va a ejecutar el defer

	fmt.Println( <-ch )

	ch1 := make(chan int)

	go func(){
		defer close(ch1)
		ch1<-1
		ch1<-2
		ch1<-3
		ch1<-4
		ch1<-5
	}()

	for n:= range ch1{
		fmt.Printf("%d\n",n)
	}

	// CREACION DE UN CANAL CON SOLO DOS ESPACIOS
	ch2 := make(chan int , 2)
	ch2<- 1
	ch2<- 2
	//ch2<- 3

	// CUANDO TU MANEJAS CANALES CON UN LEN ESTATICO PUEDES ESCRIBIR MAS DATOS SIEMPRE Y CUANDO LIBERES ESPACIO
	fmt.Println(<-ch2)
	fmt.Println(<-ch2)
	ch2<- 3
	fmt.Println(<-ch2)


}