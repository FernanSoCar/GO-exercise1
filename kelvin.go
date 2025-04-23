package main

import (
	"fmt"
)

const boilingK = 373.0

func main() {
	tempK := boilingK
	tempC := tempK - 273.0
	fmt.Printf("The boiling temperature in Kelvin is %g and in Celsius is %g.", tempK, tempC)

}