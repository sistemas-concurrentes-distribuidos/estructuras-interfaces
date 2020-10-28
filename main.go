package main

import (
	"fmt"

	"github.com/sistemas-concurrentes-distribuidos/estructuras-interfaces/model"
)

func main() {

	var contenidoWeb model.ContenidoWeb

	op := 0

	for op != 5 {
		fmt.Println("Estructuras e Interfaces")
		fmt.Println("1) Agregar Video")
		fmt.Println("2) Agregar Audio")
		fmt.Println("3) Agregar Imagen")
		fmt.Println("4) Mostrat")
		fmt.Println("5) Salir")
		fmt.Scanln(&op)

		switch op {
		case 1:
			{
				var video model.Video
				fmt.Println("VIDEO")
				fmt.Print("Titulo: ")
				fmt.Scanln(&video.Titulo)
				fmt.Print("Formato: ")
				fmt.Scanln(&video.Formato)
				fmt.Print("Fames: ")
				fmt.Scanln(&video.Frames)
				contenidoWeb.Multimedia = append(contenidoWeb.Multimedia, video)
			}
		case 2:
			{
				var audio model.Audio
				fmt.Println("AUDIO")
				fmt.Print("Titulo: ")
				fmt.Scanln(&audio.Titulo)
				fmt.Print("Formato: ")
				fmt.Scanln(&audio.Formato)
				fmt.Print("Duración: ")
				fmt.Scanln(&audio.Duracion)
				contenidoWeb.Multimedia = append(contenidoWeb.Multimedia, audio)
			}
		case 3:
			{
				var imagen model.Imagen
				fmt.Println("AUDIO")
				fmt.Print("Titulo: ")
				fmt.Scanln(&imagen.Titulo)
				fmt.Print("Formato: ")
				fmt.Scanln(&imagen.Formato)
				fmt.Print("Canales: ")
				fmt.Scanln(&imagen.Canales)
				contenidoWeb.Multimedia = append(contenidoWeb.Multimedia, imagen)
			}
		case 4:
			{
				for _, contenido := range contenidoWeb.Multimedia {
					contenido.Mostrar()
					fmt.Println("--------------------")
				}
			}
		}
	}
}
