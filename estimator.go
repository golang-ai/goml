package goml

// Estimator type for Training and Predicting 2-dimensional matrix values
type Estimator interface {
	Train([][]float64, [][]float64)
	Predict() float64
}
