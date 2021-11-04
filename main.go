package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

type problema struct {
	pregunta  string
	respuesta string
}

func main() {
	leerCSV()
}

func errorFatal(mensaje string, e error) {
	if e != nil {
		log.Fatalf(mensaje, ":", e)
	}
}

func leerCSV() {
	f, err := os.Open("problemas.csv")
	errorFatal("No se pudo leer el archivo", err)
	defer f.Close()
	csvLector := csv.NewReader(f)
	lineas, err := csvLector.ReadAll()
	errorFatal("No se pudo leer el archivo", err)
	problemas := separarLineas(lineas)

	timer := time.NewTimer(30 * time.Second)
	fmt.Printf("Tenes 30 segundos para resolver todos los problemas:\n")
	correctas := 0
	pepe := true
	for i, p := range problemas {
		if !pepe {
			break
		}
		select {
		case <-timer.C:
			fmt.Printf("Se termino el tiempo\n")
			pepe = false
		default:
			fmt.Printf("Problema numero %d: %s = \n", i+1, p.pregunta)
			var resp string
			fmt.Scanf("%s\n", &resp)
			if resp == p.respuesta {
				fmt.Println("Respuesta Correcta")
				correctas++
			} else {
				fmt.Println("Respuesta Incorrecta")
			}
		}

	}

	fmt.Printf("Respondiste %d correctamente", correctas)
}

func separarLineas(lineas [][]string) []problema {
	result := make([]problema, len(lineas))
	for i, linea := range lineas {
		result[i] = problema{
			pregunta:  linea[0],
			respuesta: linea[1],
		}
	}

	return result
}
