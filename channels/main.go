package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

// Los canales sirven para el paso de mensajes entre rutinas de go

func main() {
	fmt.Println("-- Ejercicio con channels --")
	var opcion int

	fmt.Println("Elija un ejericio: ")
	fmt.Println("1) Ejercicio con channels sencillo")
	fmt.Println("2) Ejercicio con channels avanzado")
	fmt.Scanf("%d", &opcion)

	switch opcion {
	case 1:
		fmt.Println("Has seleccionado la opcion 1")

		// Pasaremos un entero por el canal
		ch := make(chan int)
		wg.Add(2)
		fmt.Println("Comienza el programa de paso de mensajes...")

		// Receptor que imprime el mensaje
		go func() {
			i := <-ch
			fmt.Printf("Número recibido: %d\n", i)
			wg.Done()
		}()

		// Emisor del entero
		go func() {
			i := 42
			ch <- i
			fmt.Println("Hemos enviado un número")
			i = 20 // Este número no se envía
			wg.Done()
		}()

		wg.Wait()

	case 2:
		fmt.Println("Has seleccionado la opcion 2")

		// Creamos un receptor y un emisor más robusto
		ch := make(chan int)
		wg.Add(2)

		// SOLO puede ser receptor
		go func(ch <-chan int) {
			i := <-ch
			fmt.Printf("Número recibido: %d\n", i)
			wg.Done()
		}(ch)

		// SOLO puede ser emisor
		go func(ch chan<- int) {
			ch <- 42
			fmt.Println("Hemos enviado un número")
			wg.Done()
		}(ch)
		wg.Wait()
	}

}
