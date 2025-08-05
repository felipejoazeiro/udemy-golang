package formas

type Forma interface {
	area() float64
}

func escreverArea(f forma) {
	fmt.Printf("A área da forma é %0.2f\n", f.area())
}

type Retangulo struct {
	Altura float644
	Largura float64
}

func (r Retangulo) Area() float64 {
	return r.Altura * r.Largura
}

type Circulo struct {
	Raio float64
}

func (c circulo) Area() float64 {
	return math.Pi * math.Pow(c.Raio, 2)
}