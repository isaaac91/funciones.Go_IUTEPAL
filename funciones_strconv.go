package main

import (
	"fmt"
	"strconv"
)

func main() {

	valor := int64(556)
	valor2 := int64(100)
	valor3 := uint64(938109)

	//formatInt
	valInt64 := strconv.FormatInt(valor, 10)
	fmt.Printf("El valor: %d, ahora es \"%s\" tipo %T\n", valor2, valInt64, valInt64)

	newValInt64 := strconv.FormatInt(valor2, 16)
	fmt.Printf("El valor: %d, ahora es \"%s\" tipo %T\n", valor2, newValInt64, newValInt64)

	//strconv.FormatUint
	uInt64 := strconv.FormatUint(valor3, 10)
	fmt.Printf("El valor: %d, ahora es \"%s\" tipo %T\n", valor3, uInt64, uInt64)

	//strconv.ParseFloat
	valor4 := "1.55"
	valFloat, x := strconv.ParseFloat(valor4, 64)
	if x != nil {
		fmt.Println("Error en ParseFloat:", x)
	} else {
		fmt.Printf("EL valor: \"%s\" ahora es %f tipo %T\n", valor4, valFloat, valFloat)
	}

	//strconv.FormatBool
	valor5 := true
	valBool := strconv.FormatBool(valor5)
	fmt.Printf("El valor %t = \"%s\", tipo %T\n", valor5, valBool, valBool)

	//strconv.Quote
	valor6 := `quinto año "A"`
	valComillas := strconv.Quote(valor6)
	fmt.Println("Texto original:", valor6, "; Texto nuevo:", valComillas)

}
