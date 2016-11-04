package mario_go

type Dot struct {
	X     int
	Y     int
	Color int32
}

type Dots []Dot

func NewDot(x, y int, color int32) *Dot {
	d := &Dot{x, y, color}
	return d
}

func MakeDots(indices []int32) Dots {
	dots := Dots{}
	for i := 0; i < len(indices); i += 3 {
		dots = append(dots, *NewDot(int(indices[i]), int(indices[i+1]), indices[i+2]))
	}
	return dots
}
