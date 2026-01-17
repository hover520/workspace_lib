package workspace_lib

// AddNums add two numbers
//
// More information on [mathsisfun]
//
// [mathsisfun]: https://www.mathsisfun.com/numbers/addition.html

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

func AddNums[T Number](a, b T) T {
	return a + b
}

func SubNums[T Number](a, b T) T {
	return a - b
}
