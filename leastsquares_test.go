package goml_test

import "testing"

var testPredict = []struct {
	sample  [][]float64
	targets []float64
}{
	{[][]float64{{60}, {61}, {62}, {63}, {65}}, []float64{3.1, 3.6, 3.8, 4, 4.1}},
	{[][]float64{{23.6}, {28.2}, {31.3}, {33.5}, {37.1}}, []float64{1.2, 2.1, 3.8, 4.5, 5.3}},
}

func TestTrain(t *testing.T) {
	for _, v := range testPredict {

	}
}
