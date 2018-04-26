package main

import (
	"awesomeProject/alg"
	"awesomeProject/constants"
	obj_func "awesomeProject/obj_func"
	"fmt"
	"github.com/fogleman/gg"
	"gonum.org/v1/gonum/diff/fd"
	"math"
	"math/rand"
	"time"
	"log"
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func main() {
	N := 400
	xInit := make([]float64, N)
	rGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range xInit {
		xInit[i] = rGenerator.Float64() * 100
	}

	defer timeTrack(time.Now(), "crystal")

	h := constants.H
	x := make([]float64, N)
	z := make([]float64, N)
	copy(x, xInit)
	FX := obj_func.U(xInit)
	for h > constants.E {
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
		fmt.Println("h:", h)
		fmt.Println("FX", FX)
		fmt.Println("-------")
	}

	fmt.Println(x) // Print the result

	max := func(s []float64) float64 {
		m := math.Abs(s[0])
		for i, val := range s {
			if math.Abs(val) > m {
				m = math.Abs(val)
			}
			s[i] = val / 10.0;
		}
		return m
	}(x) // Find max absolute value for proper render

	dc := gg.NewContext((int)(max*0.22), (int)(max*0.22))
	dc.SetRGB255(255, 255, 255)
	dc.Clear()
	for i := 0; i < N/2; i++ {
		dc.DrawPoint(x[i]+(float64)(dc.Width()/2.0), x[i+N/2]+(float64)(dc.Width()/2.0), (float64)(dc.Width()/250.0))
	}
	dc.SetRGB255(0, 0, 0)
	dc.Fill()
	dc.SavePNG("test.png")
}
