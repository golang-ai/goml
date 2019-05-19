package goml

type Estimator interface {
	Train()
	Predict() float64
}
