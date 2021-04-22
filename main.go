package main

import "fmt"

var data = [][][]int{}

func main() {
	// tamaño := 4
	// cantidad_fils := 4

	total_casillas := [][]int{
		{1, 1}, {1, 2}, {1, 3}, {1, 4},
		{2, 1}, {2, 2}, {2, 3}, {2, 4},
		{3, 1}, {3, 2}, {3, 3}, {3, 4},
		{4, 1}, {4, 2}, {4, 3}, {4, 4},
	}

	casillasIniciales := [][]int{
		{1, 1}, {1, 2}, {1, 3}, {1, 4},
	}

	for _, casI := range casillasIniciales {
		prueba(casI, total_casillas, [][]int{}, 1)
	}

}
func prueba(casillaI []int, casillasLibres, reinasColocadas [][]int, columnaActual int) {

	// fmt.Printf("nueva reina : %v\n", casillaI)

	casillasLimpias := limpiar(casillasLibres, casillaI)
	// fmt.Printf("limpias : %v\n", casillasLimpias)

	reisColocadas := reinasColocadas
	// make([][]int, len(reinasColocadas))
	// copy(reisColocadas, reinasColocadas)
	// fmt.Printf("reinas: %v\n", casillaI)

	reisColocadas = append(reisColocadas, casillaI)

	if len(casillasLimpias) == 0 {
		// ya no quedan casillas libres
		// TODO: debe guardar la combinacion como string
		// data = append(data, reisColocadas)
		fmt.Printf("combinacion: %v\n", reisColocadas)
	} else {
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

		// if len(casillasSiguientes) > 0 {
		for _, casS := range casillasSiguientes {
			prueba(casS, casillasLimpias, reisColocadas, columnaActual+1)
		}
		// }
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
