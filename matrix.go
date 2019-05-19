package goml

import "fmt"

func validate(slice [][]float64) (bool, error) {
	cols := len(slice[0])

	for _, row := range slice {
		if len(row) != cols {
			return false, fmt.Errorf("Matrix dimensions did not match")
		}
	}

	return true, nil
}

func multiply(matrix1, matrix2 [][]float64) ([][]float64, error) {
	if len(matrix1[0]) != len(matrix2) {
		return matrix1, fmt.Errorf("Inconsistent matrix supplied")
	}

	var product [][]float64
	for k, row := range matrix1 {
		for i, colVal := range row {
			product[k][i] += colVal * matrix2[i][k]
		}
	}

	return product, nil
}

func isSquare(slice [][]float64) bool {
	cols := len(slice[0])
	rows := len(slice)

	return rows == cols
}
