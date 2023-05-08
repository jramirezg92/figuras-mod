package figuras

import "fmt"

type Geometria interface {
	calcularArea() float64
	calcularPerimetro() float64
}

func Medidas(g Geometria) {
	fmt.Println(g)
	fmt.Println(g.calcularArea())
	fmt.Println(g.calcularPerimetro())
}
