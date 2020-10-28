package model

import "fmt"

type Video struct {
	Titulo  string
	Formato string
	Frames  int
}

func (t Video) Mostrar() {
	fmt.Println("Titulo:", t.Titulo)
	fmt.Println("Formato:", t.Formato)
	fmt.Println("Frames:", t.Frames)
}
