package main

import "fmt"

type Persona struct {
	nombre   string
	estatura float64
}

func main() {
	p := Persona{"Rafael", 1.83}
	fmt.Println(p.nombre, p.correr())
}

func (p *Persona) correr() string {
	return "corriendo"
}
