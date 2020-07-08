package test

import (
	"fmt"
	"math"
	"testing"
)

func rotateInt(matrix [][]int, turn int) {
	if turn == 0 {
		return
	}
	n := len(matrix)
	x := int(math.Floor(float64(n) / 2))
	y := n - 1
	for i := 0; i < x; i++ {
		for j := i; j < y-i; j++ {
			switch turn {
			case 0: // pass
			case 1:
				matrix[i][j], matrix[j][y-i], matrix[y-j][i], matrix[y-i][y-j] = matrix[y-j][i], matrix[i][j], matrix[y-i][y-j], matrix[j][y-i]
			case 2:
				matrix[i][j], matrix[y-i][y-j] = matrix[y-i][y-j], matrix[i][j]
				matrix[j][y-i], matrix[y-j][i] = matrix[y-j][i], matrix[j][y-i]
			case 3:
				matrix[i][j], matrix[j][y-i], matrix[y-j][i], matrix[y-i][y-j] = matrix[j][y-i], matrix[y-i][y-j], matrix[i][j], matrix[y-j][i]
			}
		}
	}
}

func printInt(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%v ", matrix[i][j])
		}
		fmt.Println()
	}
}

func TestRotate(t *testing.T) {
	m := [][]int{{1, 2, 3, 4}, {12, 13, 14, 5}, {11, 16, 15, 6}, {10, 9, 8, 7}}
	printInt(m)
	rotateInt(m, 3)
	fmt.Println("----------------")
	printInt(m)
}
