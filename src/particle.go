package src

import (
	"log"
)

type Particle interface {
	Advance()
}

type particle struct {
	mass, radius float64
	Pos          Position
	Vel          Velocity
	Acc          Acceleration
}

func (p *particle) Advance() {

	log.Println("Advancing 1 step in time...", p.Pos.X, p.Pos.Y, p.Pos.Z)
}

type Position struct {
	X, Y, Z float64
}

type Velocity struct {
	X, Y, Z float64
}

type Acceleration struct {
	X, Y, Z float64
}

func NewParticle(initPos Position, initVel Velocity, initAcc Acceleration) Particle {
	return &particle{
		Pos: initPos,
		Vel: initVel,
		Acc: initAcc,
	}
}
