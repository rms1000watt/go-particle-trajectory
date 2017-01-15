package src

import "log"

func Freefall(cfg Config) {
	// Init phase
	dt := 0.1
	dtVector := ToVector(dt)
	g := NewVector(0, -9.8, 0)

	p := &particle{
		Pos: NewVector(1, 1, 1),
		Acc: g,
	}

	// Simulate phase
	for i := 0; i < cfg.Iterations; i++ {
		p.Vel = SumVectors(p.Vel, MultiplyVectors(p.Acc, dtVector))
		p.Pos = SumVectors(p.Pos, MultiplyVectors(p.Vel, dtVector))
		p.PosArr = append(p.PosArr, p.Pos)

		log.Println(i, "Vel", p.Vel)
		log.Println(i, "Pos", p.Pos)
	}

	// TODO: Plot phase
}
