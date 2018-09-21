package main

import (
	"awesomeProject/alg"
	"awesomeProject/constants"
	"awesomeProject/obj_func"
	"fmt"
	"github.com/fogleman/gg"
	"gonum.org/v1/gonum/diff/fd"
	"math"
	"math/rand"
	"time"
)

func main() {
	startFull := time.Now()
	N := 400
	xInit := make([]float64, N)
	rGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range xInit {
		xInit[i] = rGenerator.Float64() - 0.5

		xInit[i] *= 1500
	}

	printDots(xInit, N)

	h := constants.H
	x := make([]float64, N)
	z := make([]float64, N)
	copy(x, xInit)
	FX := obj_func.U(xInit)
	for h > constants.E {
		start := time.Now()
		G := fd.Gradient(nil, obj_func.U, x, nil) // Gradient vector
		V := alg.Unit_vec(G)                      // Unit gradient vector
		z = alg.Sub_vec(x, alg.Mult_vec(V, h))
		FZ := obj_func.U(z)
		if FZ < FX {
			x = z
			FX = FZ
		} else {
			if h >= constants.E {
				h = h / 1.1
			} else {
				break
			}
		}
		elapsed := time.Since(start)
		fmt.Printf("step took %s\n", elapsed)
		fmt.Println("h:", h)
		fmt.Println("FX", FX)
		fmt.Println("-------")
	}
	elapsedFull := time.Since(startFull)
	fmt.Printf("Full took %s\n", elapsedFull)
	fmt.Println(x) // Print the result

	printDots(x, N)
}

func printDots(x []float64, N int) {
	k := func(s []float64) float64 {
		m := math.Abs(s[0])
		for _, val := range s {
			if math.Abs(val) > m {
				m = math.Abs(val)
			}
		}
		return m
	}(x)
	// Find max absolute value for proper render
	k = 750 / k
	dc := gg.NewContext(1500, 1500)
	dc.SetRGB255(255, 255, 255)
	dc.Clear()
	for i := 0; i < N/2; i++ {
		dc.DrawPoint(x[i]*k+float64(dc.Width()/2), x[i+N/2]*k+float64(dc.Width()/2), (float64)(dc.Width()/200.0))
	}
	dc.SetRGB255(255, 0, 0)
	dc.DrawPoint(constants.POS_X_1*k+float64(dc.Width()/2), constants.POS_Y_1*k+float64(dc.Width()/2), float64(dc.Width()/100.0))
	dc.DrawPoint(constants.POS_X_2*k+float64(dc.Width()/2), constants.POS_Y_2*k+float64(dc.Width()/2), float64(dc.Width()/100.0))
	dc.SetRGB255(0, 0, 0)
	dc.Fill()
	dc.SavePNG("200points_1000.png")
}
