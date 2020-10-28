package model

import "fmt"

type Imagen struct {
	Titulo  string
	Formato string
	Canales string
}

func (t Imagen) Mostrar() {
	fmt.Println("Titulo:", t.Titulo)
	fmt.Println("Formato:", t.Formato)
	fmt.Println("Canales:", t.Canales)
}
