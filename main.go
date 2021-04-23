package main

import (
	"fmt"
	"sync"
)

var data = []string{}
var tamaño = 0
var iteraciones int = 0
var combinaciones int = 0

var wg sync.WaitGroup

func main() {
	tamaño = 8
	total_casillas := [][]int{}

	for i := 1; i <= tamaño; i++ {
		for j := 1; j <= tamaño; j++ {
			aux := []int{i, j}
			total_casillas = append(total_casillas, aux)
		}
	}

	primerAnalisi(tamaño, total_casillas)

	// time.Sleep(2 * time.Second)

	fmt.Println("iteraciones finales ", iteraciones)
	fmt.Println(combinaciones)

}

func primerAnalisi(tamaño int, casillasTotales [][]int) {
	// fmt.Println("iniciando analisis")
	casillasIniciales := [][]int{}
	for _, casI := range casillasTotales {
		if casI[0] != 1 {
			break
		}
		casillasIniciales = append(casillasIniciales, casI)
	}

	iterChan := make(chan bool)

	// wg.Add(1)
	go func() {
		iterNums := 0
		for {
			if <-iterChan {
				fmt.Println("comb")
				iterNums++
				// if iterNums == iteraciones {
				// 	wg.Done()
				// 	break
				// }
			}
		}
	}()

	// aqui empieza el analisis
	// iteraciones = iteraciones + len(casillasIniciales)
	wg.Add(len(casillasIniciales))
	for _, casI := range casillasIniciales {
		iteraciones++
		// wg.Add(1)
		go analizar(casI, casillasTotales, [][]int{}, 1, iterChan)
	}
	wg.Wait()

}

func analizar(casillaI []int, casillasLibres, reinasColocadas [][]int,
	columnaActual int, iterChan chan bool) {

	// fmt.Printf("nueva reina : %v\n", casillaI)
	casillasLimpias := limpiar(casillasLibres, casillaI)
	// fmt.Printf("limpias : %v\n", casillasLimpias)
	reisColocadas := reinasColocadas
	// fmt.Printf("reinas: %v\n", casillaI)
	reisColocadas = append(reisColocadas, casillaI)

	if len(casillasLimpias) == 0 {
		// ya no quedan casillas libres
		// data = append(data, reisColocadas)

		if len(reisColocadas) == tamaño {
			// aui encontro una combinacion valida
			// TODO: debe guardar la combinacion como string
			fmt.Printf("combinacion: %v\n", reisColocadas)
			combinaciones++
			iterChan <- true
			wg.Done()
		} else {
			wg.Done()
		}
	}
	if len(casillasLimpias) > 0 {
		// todavia quedan casillas a analizar
		casillasSiguientes := [][]int{}

		for _, casillaL := range casillasLimpias {
			if casillaL[0] == columnaActual+1 {
				// pertenece a la columna siguiente
				casillasSiguientes = append(casillasSiguientes, casillaL)
			}
			if casillaL[0] != columnaActual+1 {
				break
			}
		}

		if len(casillasSiguientes) > 0 {
			// iteraciones = iteraciones + len(casillasSiguientes)
			wg.Add(len(casillasSiguientes))
			for _, casS := range casillasSiguientes {
				iteraciones++
				go analizar(casS, casillasLimpias, reisColocadas, columnaActual+1, iterChan)
			}
			wg.Done()
		} else {
			// ya no hay casillas de la columna siguiente
			// y no hay nada que hacer
			iterChan <- true
			wg.Done()

		}

	}
}

func limpiar(casillasLibres [][]int, nuevaReina []int) [][]int {

	columnaOcupada, filaOcupada := nuevaReina[0], nuevaReina[1]
	crecienteOcupada := columnaOcupada + filaOcupada
	decrecienteOcupada := columnaOcupada - filaOcupada

	nuevasCasillasLibres := [][]int{}

	for i := 0; i < len(casillasLibres); i++ {
		casilla := casillasLibres[i]

		if casilla[0] != columnaOcupada && casilla[1] != filaOcupada {
			// no esta ni en la misma fila ni columna

			creciente := casilla[0] + casilla[1]
			decreciente := casilla[0] - casilla[1]
			if creciente != crecienteOcupada && decreciente != decrecienteOcupada {
				// no esta en las mismas diagonales
				nuevasCasillasLibres = append(nuevasCasillasLibres, casilla)
			}
		}
	}

	return nuevasCasillasLibres
}
