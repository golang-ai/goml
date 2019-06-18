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
	for k := range samples {
		e.Targets[k] = append(e.Targets[k], targets...)
	}

	e.computeCoefficients()
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
		fmt.Println(k, v, sample[k], result)
		result += v * sample[k]
	}
	fmt.Println(result)
	return result
}

// coefficient(b) = (X'X)-1X'Y
func (e *Regression) computeCoefficients() error {
	samplesMatrix := e.getSamplesMatrix()
	targetsMatrix := e.getTargetsMatrix()
	// fmt.Println(samplesMatrix.samples)
	ts, err := samplesMatrix.transpose().multiply(samplesMatrix.samples)
	tf, er := samplesMatrix.transpose().multiply(targetsMatrix.samples)

	if err != nil {
		return err
	}

	ts.inverse()

	if er != nil {
		return er
	}
	fmt.Println(ts, tf)
	// already checked squared matrix
	m, _ := ts.multiply(tf.samples)

	e.coefficients = m.getColumnValues(0)
	e.intercept = e.coefficients[0]
	e.coefficients = e.coefficients[1:]

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
