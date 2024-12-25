package algebra

import "errors"

func Mult_matrix(matrix1, matrix2 [][]int) ([][]int, error) {

	if len(matrix1[0]) != len(matrix2) {
		return nil, errors.New("Não é possível multiplicar as matrizes")

	}
	rows := len(matrix1)
	cols := len(matrix2[0])

	result := make([][]int, rows)

	for i := range result {
		result[i] = make([]int, cols)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			for k := 0; k < len(matrix1[0]); k++ {

				result[i][j] += matrix1[i][k] * matrix2[k][j]
			}
		}
	}

	return result, nil

}
