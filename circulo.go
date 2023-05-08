package figuras

import "math"

type Circulo struct {
	Radio float64
}

func (cir *Circulo) calcularArea() float64 {
	pi := math.Pi
	return pi * (cir.Radio * cir.Radio)
}

func (cir *Circulo) calcularPerimetro() float64 {
	pi := math.Pi
	return 2 * pi * cir.Radio
}
