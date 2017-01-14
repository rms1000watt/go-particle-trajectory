package src

type World interface {
	Advance() error
	AddParticles([]Particle)
}

type world struct {
	DT        float64
	Acc       Acceleration
	Boundary  Boundary
	Particles []Particle
}

func (w *world) Advance() error {
	for _, particle := range w.Particles {
		particle.Advance()
	}
	return nil
}

func (w *world) AddParticles(particles []Particle) {
	w.Particles = particles
}

type Boundary struct {
}

func NewWorld() World {
	return &world{}
}
