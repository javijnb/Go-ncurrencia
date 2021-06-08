package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Para sincronizar rutinas de go:
var wg = sync.WaitGroup{}

// Read-write mutex
var m = sync.RWMutex{}

var counter = 0

func main() {
	fmt.Println("-- GO-NCURRENCIA --")

	runtime.GOMAXPROCS(100)
	fmt.Printf("Hilos: %d\n", runtime.GOMAXPROCS(-1))

	for i := 0; i < 10; i++ {

		// Indicamos cuantas rutinas go añadimos al main (sayHello + increment = 2)
		wg.Add(2)

		// Mutex de lectura, no cambiamos información, imprimimos mensaje
		m.RLock()
		go sayHello()

		// Mutex de escritura, alteramos el valor del contador
		m.Lock()
		go increment()

	}

	// Indicar cuándo ha finalizado la ejecución de rutinas (fase liberación) hasta que contador de WG sea 0
	// Cuando el contador es zero continúa más allá de la línea 36

	wg.Wait()

	fmt.Println("Hemos acabado con la ejecución del programa")

}

func sayHello() {
	fmt.Printf("Hola #%d\n", counter)

	// Liberamos lectura
	m.RUnlock()

	// Indicamos al WaitGroup cuándo hemos acabado con esta rutina (decrementa WG en 1)
	wg.Done()
}

func increment() {
	counter++

	// Liberamos escritura
	m.Unlock()

	// Indicamos al WaitGroup cuándo hemos acabado con esta rutina (decrementa WG en 1)
	wg.Done()
}
