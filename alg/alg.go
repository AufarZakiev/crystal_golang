package alg

import (
	"math"
)

func Unit_vec(vec []float64) []float64 { // Unit vector computation vector
	vec_len := func() float64 { // Vector length compuation function
		var length float64
		for _, x := range vec {
			length += x * x
		}
		return math.Sqrt(length)
	}()
	for i, x := range vec {
		vec[i] = x / vec_len
	}
	return vec
}

func Mult_vec(vec []float64, multi float64) []float64 {
	for i, val := range vec {
		vec[i] = val * multi
	}
	return vec
}

func Add_vec(vec []float64, add []float64) []float64 {
	for i, val := range vec {
		vec[i] = val + add[i]
	}
	return vec
}

func Sub_vec(vec []float64, sub []float64) []float64 {
	for i, val := range vec {
		vec[i] = val - sub[i]
	}
	return vec
}
