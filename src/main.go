package main

import (
	"fmt"
	"studies/algebra"
)

func main() {

	matrix1 := [][]int{
		{1, 2},
		{3, 4},
		{4, 3},
	}

	matrix2 := [][]int{
		{1, 2},
		{3, 4},
		{5, 6}}

	resultsum_matrix, err := algebra.Sum_matrix(matrix1, matrix2)

	//result1 := algebra.TransposeMatrix(matrix1)
	//resutult2 := algebra.TransposeMatrix(matrix2)

	fmt.Printf("%d - %s\n", resultsum_matrix, err)

}
