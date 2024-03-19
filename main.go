package main

import "fmt"

func main() {
	var num int
	fmt.Println("1-consultar tarifa")
	fmt.Println("2-cobro de peaje")
	fmt.Println("3-cambiar tarifa")
	fmt.Println("4-mostras tarifas")
	fmt.Println("5-salir")
	fmt.Scanln(&num)
}
