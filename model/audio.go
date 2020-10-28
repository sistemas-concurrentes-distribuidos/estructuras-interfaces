package model

import "fmt"

type Audio struct {
	Titulo   string
	Formato  string
	Duracion float64
}

func (t Audio) Mostrar() {
	fmt.Println("Titulo:", t.Titulo)
	fmt.Println("Formato:", t.Formato)
	fmt.Println("Duracion:", t.Duracion)
}
