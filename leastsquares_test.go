package goml_test

import (
	"fmt"
	"goml"
	"testing"
)

var testPredict = []struct {
	sample  [][]float64
	targets []float64
	predict [][]float64
	result  float64
}{
	{[][]float64{{60}, {61}, {62}, {63}, {65}}, []float64{3.1, 3.6, 3.8, 4, 4.1}, [][]float64{{64}}, 4.06},
	{[][]float64{{23.6}, {28.2}, {31.3}, {33.5}, {37.1}}, []float64{1.2, 2.1, 3.8, 4.5, 5.3}, [][]float64{{35.5}}, 4.9},
}

func TestTrain(t *testing.T) {
	e := &goml.Regression{}
	for _, object := range testPredict {
		e.Train(object.sample, object.targets)
		e.Predict(object.predict)

		if e.Predicted[0] != object.result {
			t.Fatalf("want %f, got %f", object.result, e.Predicted[0])
		}

		fmt.Println(e.Predicted)
	}
}
