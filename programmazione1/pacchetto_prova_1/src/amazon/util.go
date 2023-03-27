package amazon

import (
	"fmt"
	"math"
)

func (x Box) Volume() int{
	return x.L * x.H * x.D
}

func (x Box) Length() int {
	return x.L
}

func (x Box) String() string {
	return fmt.Sprintf("Lunghezza: %d, Altezza %d, Profondità: %d", x.L, x.H, x.D)
}

func (x Sphere) Radius() int {
	return x.R
}

func (x Sphere) Volume() int {
	return int((4.0 / 3.0) * math.Pi * float64(x.R*x.R*x.R))
}

func (x Sphere) String() string {
	return fmt.Sprintf("radius: %d", x.R)
}