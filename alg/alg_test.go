package alg

import (
	"reflect"
	"testing"
)

func TestAdd_vec(t *testing.T) {
	type testpair struct {
		vec1 []float64
		vec2 []float64
		sum  []float64
	}
	var tests = []testpair{
		{[]float64{1, 1, 1}, []float64{1, 1, 1}, []float64{2, 2, 2}},
		{[]float64{0, 0, 0}, []float64{0, 0, 0}, []float64{0, 0, 0}},
		{[]float64{1, 2, 3, 4}, []float64{4, 3, 2, 1}, []float64{5, 5, 5, 5}},
		{[]float64{99, 1, 50, 99, 1, 50}, []float64{1, 1, 1, 1, 1, 1}, []float64{100, 2, 51, 100, 2, 51}},
		{[]float64{99.99, 1.99, 50.99, 99.99, 1.99, 50.99}, []float64{0.01, 0.01, 0.01, 0.01, 0.01, 0.01}, []float64{100, 2, 51, 100, 2, 51}},
	}
	for _, test := range tests {
		v := Add_vec(test.vec1, test.vec2)
		if !reflect.DeepEqual(v, test.sum) {
			t.Error(
				"For values", test.vec1, test.vec2,
				"expected", test.sum,
				"got", v,
			)
		}
	}
}

func TestSub_vec(t *testing.T) {
	type testpair struct {
		vec1 []float64
		vec2 []float64
		sub  []float64
	}
	var tests = []testpair{
		{[]float64{1, 1, 1}, []float64{1, 1, 1}, []float64{0, 0, 0}},
		{[]float64{0, 0, 0}, []float64{0, 0, 0}, []float64{0, 0, 0}},
		{[]float64{1, 2, 3, 4}, []float64{4, 3, 2, 1}, []float64{-3, -1, 1, 3}},
		{[]float64{99, 1, 50, 99, 1, 50}, []float64{1, 1, 1, 1, 1, 1}, []float64{98, 0, 49, 98, 0, 49}},
		{[]float64{100, 2, 51, 100, 2, 51}, []float64{0.01, 0.01, 0.01, 0.01, 0.01, 0.01}, []float64{99.99, 1.99, 50.99, 99.99, 1.99, 50.99}},
	}
	for _, test := range tests {
		v := Sub_vec(test.vec1, test.vec2)
		if !reflect.DeepEqual(v, test.sub) {
			t.Error(
				"For values", test.vec1, test.vec2,
				"expected", test.sub,
				"got", v,
			)
		}
	}
}

func TestMult_vec(t *testing.T) {
	type testpair struct {
		vec1    []float64
		multi   float64
		product []float64
	}
	var tests = []testpair{
		{[]float64{1, 1, 1}, 3, []float64{3, 3, 3}},
		{[]float64{2, 2, 2}, 0.5, []float64{1, 1, 1}},
		{[]float64{34.6, 54.1, 21.7}, 10, []float64{346, 541, 217}},
		{[]float64{36.6, 54.1, 21.7}, 0.1, []float64{3.66, 5.41, 2.17}},
	}
	for _, test := range tests {
		v := Mult_vec(test.vec1, test.multi)
		if !reflect.DeepEqual(v, test.product) {
			t.Error(
				"For values", test.vec1, test.multi,
				"expected", test.product,
				"got", v,
			)
		}
	}
}

func TestUnit_vec(t *testing.T) {
	type testpair struct {
		vec      []float64
		unit_vec []float64
	}
	var tests = []testpair{
		{[]float64{1, 1, 1, 1}, []float64{0.5, 0.5, 0.5, 0.5}},
		{[]float64{5, 5, 5, 5}, []float64{0.5, 0.5, 0.5, 0.5}},
	}
	for _, test := range tests {
		v := Unit_vec(test.vec)
		if !reflect.DeepEqual(v, test.unit_vec) {
			t.Error(
				"For values", test.vec,
				"expected", test.unit_vec,
				"got", v,
			)
		}
	}
}
