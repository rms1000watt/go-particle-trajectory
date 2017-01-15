package src

func NewVector(x, y, z float64) [3]float64 {
	return [3]float64{x, y, z}
}

func ToVector(v float64) [3]float64 {
	return NewVector(v, v, v)
}

func SumVectors(vectors ...[3]float64) [3]float64 {
	out := [3]float64{}
	for _, vector := range vectors {
		out[0] += vector[0]
		out[1] += vector[1]
		out[2] += vector[2]
	}
	return out
}

func MultiplyVectors(vectors ...[3]float64) [3]float64 {
	out := [3]float64{1, 1, 1}
	for _, vector := range vectors {
		out[0] *= vector[0]
		out[1] *= vector[1]
		out[2] *= vector[2]
	}
	return out
}
