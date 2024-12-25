package algebra

import (
	"errors"
)

func Sum_matrix(matrix1, matrix2 [][]int) ([][]int, error) {

	if len(matrix1) != len(matrix2) || len(matrix1[0]) != len(matrix2[0]) {
		return nil, errors.New("Não é possível somar: as matrizes não têm as mesma proporção.")

	}
	rows := len(matrix1)
	cols := len(matrix2[0])

	sum_result := make([][]int, rows)

	for i := range sum_result {
		sum_result[i] = make([]int, cols)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {

			sum_result[i][j] = matrix1[i][j] + matrix2[i][j]
		}
	}

	return sum_result, nil

}
