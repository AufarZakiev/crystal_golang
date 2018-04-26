package obj_func

import (
	"awesomeProject/constants"
	"math"
)

func U(input []float64) float64 { // Wrapper around objective function to fit into fd.Gradient function
	if len(input)%2 != 0 { // If x coordinates count are not equal to y coordinates count
		panic("Coordintaes are not 2D")
	}
	return func(x []float64, y []float64) float64 { //anonymous function to compute objective function
		if len(x) != len(y) {
			panic("Coordinates are not consistent")
		}
		N := len(x)

		summingChan1 := make(chan float64, N)
		summingChan2 := make(chan float64, N)
		summingChan3 := make(chan float64, 2)
		summingChan4 := make(chan float64, 2)

		go func() {
			var firstSum float64
			for i := 0; i < N; i++ {
				firstSum += <-summingChan1
			}
			summingChan3 <- firstSum
		}()

		go func() {
			var secondSum float64
			for i := 0; i < N; i++ {
				secondSum += <-summingChan2
			}
			summingChan4 <- secondSum
		}()

		for i := 0; i < N; i++ {

			go func(index int) {
				partial_sum := constants.A / (math.Sqrt(x[index]*x[index]+y[index]*y[index]) + constants.L1)
				summingChan1 <- partial_sum
			}(i)

			go func(index int) {
				var partial_sum float64
				for j := index + 1; j < N; j++ {
					partial_sum += constants.B / (math.Sqrt((x[index]-x[j])*(x[index]-x[j])+(y[index]-y[j])*(y[index]-y[j])) + constants.L2)
				}
				summingChan2 <- partial_sum
			}(i)
		}

		return <-summingChan4 - <-summingChan3
	}(input[:len(input)/2], input[len(input)/2:]) // return value of objective function
}
