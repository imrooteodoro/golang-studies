package algebra

func TransposeMatrix(matrix [][]int) [][]int {

	rows := len(matrix)
	cols := len(matrix[0])

	transposed_matrix := make([][]int, cols)

	for i := range transposed_matrix {

		transposed_matrix[i] = make([]int, rows)
	}

	for i := 0; i < cols; i++ {
		for j := 0; j < rows; j++ {
			transposed_matrix[j][i] = matrix[i][j]

		}
	}

	return transposed_matrix
}
