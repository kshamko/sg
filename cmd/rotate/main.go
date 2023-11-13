package main

import (
	"fmt"
)

func main() {

	testCases := []struct {
		arr      [][]int
		expected string
	}{
		{
			arr: [][]int{
				{2, 3, 4, 8},
				{5, 7, 9, 12},
				{1, 0, 6, 10},
			},
			expected: "2, 3, 4, 8, 12, 10, 6, 0, 1, 5, 7, 9",
		},
		{
			arr:      [][]int{},
			expected: "",
		},
		{
			arr: [][]int{
				{2, 3, 4, 8},
				{5, 7, 9, 12},
				{1, 0, 6, 10},
				{1, 0, 6, 10},
			},
			expected: "2, 3, 4, 8, 12, 10, 10, 6, 0, 1, 1, 5, 7, 9, 6, 0",
		},
	}

	for i, testCase := range testCases {
		res := clockwise(testCase.arr)
		if testCase.expected != res {
			panic(fmt.Sprintf("test case %d failed: %s != %s", i, testCase.expected, res))
		}

		fmt.Println(res)
	}

	fmt.Println("test cases passed")
}

/*
clockwise function makes all the magic
Algirithm:
 1. take first row of input slice and append it to the result
 2. Remove first row from input slice.
*/
func clockwise(in [][]int) string {

	if len(in) == 0 {
		return ""
	}

	res := []int{}

	for {
		res = append(res, in[0]...)

		in = in[1:]
		in = counterClock(in)

		if len(in) == 1 {
			res = append(res, in[0]...)
			break
		}
	}

	return intSliceToString(res, ", ")
}

func newArr(in [][]int) [][]int {

	numRows := len(in)
	numCols := len(in[0])

	out := make([][]int, numCols)

	for i := range out {
		out[i] = make([]int, numRows)
	}

	return out

}

/*
Intuition for counterClock:

Input:
5 7 9 12
1 0 6 10

# Rotated
12 10
9 6
7 0
5 1

# Indexes mapping
[0, 3] -> [0, 0]
[1, 3] -> [0, 1]

[0, 2] -> [1, 0]
[1, 2] -> [1, 1]

[0, 1] -> [2, 0]
[1, 1] -> [2, 1]

[0, 0] -> [3, 0]
[1, 0] -> [3, 1]
*/
func counterClock(in [][]int) [][]int {
	out := newArr(in)

	for i := range in {
		for j := range in[i] {

			newI := -1*(j-len(in[i])) - 1
			newJ := i

			out[newI][newJ] = in[i][j]
		}
	}

	return out
}

// intSliceToString is a helper function to transform []int slice into a string
func intSliceToString(in []int, delim string) string {

	var res string

	for i, num := range in {
		res += fmt.Sprintf("%d", num)
		if i < len(in)-1 {
			res += delim
		}
	}
	return res
}
