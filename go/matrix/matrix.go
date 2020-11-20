package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

type Matrix [][]int

func New(s string) (Matrix, error) {
	rowsOfStrings := strings.Split(s, "\n")
	matrixAsStrings := make([][]string, len(rowsOfStrings))
	for idx, str := range rowsOfStrings {
		matrixAsStrings[idx] = strings.Fields(str)
	}

	matrix := make(Matrix, len(rowsOfStrings))

	rowLength := -1
	for i := range matrixAsStrings {
		if rowLength < 0 {
			rowLength = len(matrixAsStrings[i])
		} else if rowLength != len(matrixAsStrings[i]) {
			return nil, fmt.Errorf("Unequal row length")
		}
		matrix[i] = make([]int, rowLength)
		for j := range matrixAsStrings[i] {
			v, err := strconv.Atoi(matrixAsStrings[i][j])
			if err != nil {
				return nil, err
			}
			matrix[i][j] = v
		}
	}

	return matrix, nil
}

func (inMat Matrix) Cols() (outMat [][]int) {
	rowLength := len(inMat[0])
	outMat = make([][]int, rowLength)
	for _, row := range inMat {
		for i, v := range row {
			outMat[i] = append(outMat[i], v)
		}
	}
	return
}

func (inMat Matrix) Rows() (outMat [][]int) {
	columnLength := len(inMat)
	rowLength := len(inMat[0])
	outMat = make([][]int, columnLength)
	data := make([]int, columnLength*rowLength)
	for i := range inMat {
		start := i * rowLength
		end := start + rowLength
		outMat[i] = data[start:end:end]
		copy(outMat[i], inMat[i])
	}
	return
}

func (mat Matrix) Set(row, col, val int) bool {
	if row < 0 || col < 0 || row >= len(mat) || col >= len(mat[0]) {
		return false
	}
	mat[row][col] = val
	return true
}
