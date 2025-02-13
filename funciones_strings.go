package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var valor string

	lectura := bufio.NewScanner(os.Stdin)
	fmt.Print("Ingrese la frase que desee:\n")
	lectura.Scan()
	valor = lectura.Text()
	vocales := "aeiouáéíóúAEIOUÁÉÍÓÚ"
	//strings.ToLower
	minuscula := strings.ToLower(valor)
	fmt.Println("La frase", valor, "en minúsculas sería:", minuscula)

	//strings.Count
	contador := strings.Count(valor, "")
	fmt.Println("La frase", valor, "contiene", contador-1, "caracteres")

	//strings.ContainsAny
	comprobarVocal := strings.ContainsAny(valor, vocales)
	if comprobarVocal == true {
		fmt.Println("La frase", valor, "SÍ contiene vocal/es")
	} else {
		fmt.Println("La frase", valor, "NO contiene vocal/es")
	}

	//strings.IndexAny
	encontrarVocal := strings.IndexAny(valor, vocales)
	fmt.Println("En la frase", valor, "la primera vocal se encuentra en el caracter #", encontrarVocal+1)

	//strings.LastIndexAny
	ultimaVocal := strings.LastIndexAny(valor, vocales)
	fmt.Println("En la frase", valor, "la última vocal se encuentra en el caracter #", ultimaVocal+1)
}
