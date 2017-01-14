package src

func Start(cfg Config) {
	iterations := 10

	particles := createParticles(10)
	world := NewWorld()
	world.AddParticles(particles)

	for i := 0; i < iterations; i++ {
		world.Advance()
	}
}

func createParticles(numberOfParticles int) []Particle {
	particles := []Particle{}
	for i := 0; i < numberOfParticles; i++ {
		particles = append(particles, &particle{})
	}
	return particles
}
