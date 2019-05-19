package goml

// Regression includes samples, targets and prediction outcome
type Regression struct {
	samples [][]float64
	targets []float64
	predict []float64
}

// Calculate process training and returns prediction
func Calculate(e Estimator) float64 {
	e.Train()
	return e.Predict()
}

func (e *Regression) Train() {

}

func (e *Regression) Predict() float64 {

}

func computeCoefficients() {

}
