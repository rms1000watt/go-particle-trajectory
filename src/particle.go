package src

type particle struct {
	mass, radius float64
	Pos          [3]float64
	Vel          [3]float64
	Acc          [3]float64
	PosArr       [][3]float64
}
