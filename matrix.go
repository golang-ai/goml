package goml

import (
	"fmt"
)

type mtx [][]float64

type matrix struct {
	samples mtx
}

func validate(slice [][]float64) (bool, error) {
	cols := len(slice[0])

	for _, row := range slice {
		if len(row) != cols {
			return false, fmt.Errorf("Matrix dimensions did not match")
		}
	}

	return true, nil
}

func (m *matrix) multiply(matrix2 [][]float64) (*matrix, error) {
	if len(m.samples[0]) != len(matrix2) {
		return &matrix{samples: m.samples}, fmt.Errorf("Inconsistent matrix supplied")
	}

	product, _ := dot(m.samples, matrix2)

	return &matrix{samples: product}, nil
}

func dot(x, y [][]float64) ([][]float64, error) {
	out := make([][]float64, len(x))

	for i := 0; i < len(x); i++ {
		out[i] = make([]float64, len(y[0]))
		for j := 0; j < len(y[0]); j++ {
			for k := 0; k < len(y); k++ {
				out[i][j] += x[i][k] * y[k][j]
			}
		}
	}

	return out, nil
}

func (m *matrix) transpose() *matrix {
	r := make([][]float64, len(m.samples[0]))
	for x := range r {
		r[x] = make([]float64, len(m.samples))
	}

	for y, s := range m.samples {
		for x, e := range s {
			r[x][y] = e
		}
	}

	return &matrix{samples: r}
}

func (m *matrix) inverse() error {
	if !isSquare(m.samples) {
		return fmt.Errorf("Matrix is not square matrix")
	}

	m.samples.lu()

	return nil
}

func isSquare(slice [][]float64) bool {
	cols := len(slice[0])
	rows := len(slice)

	return rows == cols
}

func (m *matrix) getColumnValues(column int) []float64 {
	var result []float64

	for _, v := range m.samples {
		result = append(result, v[column])
	}

	return result
}

func zero(n int) mtx {
	r := make([][]float64, n)
	a := make([]float64, n*n)
	for i := range r {
		r[i] = a[n*i : n*(i+1)]
	}
	return r
}

func eye(n int) mtx {
	r := zero(n)
	for i := range r {
		r[i][i] = 1
	}

	return r
}

func (a mtx) pivotize() mtx {
	p := eye(len(a))
	for j, r := range a {
		max := r[j]
		row := j
		for i := j; i < len(a); i++ {
			if a[i][j] > max {
				max = a[i][j]
				row = i
			}
		}
		if j != row {
			// swap rows
			p[j], p[row] = p[row], p[j]
		}
	}
	return p
}

func (m1 mtx) mul(m2 mtx) mtx {
	r := zero(len(m1))
	for i, r1 := range m1 {
		for j := range m2 {
			for k := range m1 {
				r[i][j] += r1[k] * m2[k][j]
			}
		}
	}
	return r
}

func (a mtx) lu() (l, u, p mtx) {
	l = zero(len(a))
	u = zero(len(a))
	p = a.pivotize()
	a = p.mul(a)

	for j := range a {
		l[j][j] = 1
		for i := 0; i <= j; i++ {
			sum := 0.
			for k := 0; k < i; k++ {
				sum += u[k][j] * l[i][k]
			}
			u[i][j] = a[i][j] - sum
		}

		for i := j; i < len(a); i++ {
			sum := 0.
			for k := 0; k < j; k++ {
				sum += u[k][j] * l[i][k]
			}
			l[i][j] = (a[i][j] - sum) / u[j][j]
		}
	}

	return
}
