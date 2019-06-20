package goml

import "fmt"

// Regression includes samples, targets and prediction outcome
type Regression struct {
	Samples      [][]float64
	Targets      [][]float64
	Predicted    []float64
	coefficients []float64
	intercept    float64
}

// Train datasets
func (e *Regression) Train(samples [][]float64, targets []float64) {
	e.Samples = make([][]float64, len(samples))

	// merge samples
	for k, v := range samples {
		e.Samples[k] = append(e.Samples[k], v...)
	}

	e.Targets = make([][]float64, len(targets))

	// merge targets
	for k, v := range targets {
		e.Targets[k] = append(e.Targets[k], v)
	}

	err := e.computeCoefficients()

	if err != nil {
		panic(err)
	}
}

// Predict performs prediction on samples
func (e *Regression) Predict(samples [][]float64) {
	e.Predicted = make([]float64, len(samples))

	for k, v := range samples {
		e.Predicted[k] = e.PredictSample(v)
	}
}

// PredictSample predicts on samples
func (e *Regression) PredictSample(sample []float64) float64 {
	result := e.intercept

	for k, v := range e.coefficients {
		result += v * sample[k]
	}

	return result
}

// coefficient(b) = (X'X)-1X'Y
func (e *Regression) computeCoefficients() error {
	samplesMatrix := e.getSamplesMatrix()
	targetsMatrix := e.getTargetsMatrix()

	fmt.Println(samplesMatrix.samples, targetsMatrix.samples)
	ts, err := samplesMatrix.transpose().multiply(samplesMatrix.samples)
	tf, er := samplesMatrix.transpose().multiply(targetsMatrix.samples)

	if err != nil {
		return err
	}

	l, _ := ts.inverse()

	if er != nil {
		return er
	}

	// already checked squared matrix
	ts.samples = l
	fmt.Println("ts, tf", ts.samples, tf.samples)
	m, _ := ts.multiply(tf.samples)

	fmt.Println("ms; ", m.samples)
	e.coefficients = m.getColumnValues(0)
	//fmt.Println(e.coefficients)

	e.intercept = e.coefficients[0]
	e.coefficients = e.coefficients[1:]
	fmt.Println("intercept and coeff: ", e.intercept, e.coefficients)

	return nil
}

func (e *Regression) getSamplesMatrix() *matrix {
	samples := make([][]float64, len(e.Samples))

	for k, v := range e.Samples {
		samples[k] = append([]float64{1}, v...)
	}

	return &matrix{samples: samples}
}

func (e *Regression) getTargetsMatrix() *matrix {
	return &matrix{samples: e.Targets}
}
