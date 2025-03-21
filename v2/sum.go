// sum is a stubby module that adds some stuff
package sum

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer
	constratins.Float
}

// Add adds two numbers and returns their sum. See [MathIsFun]
//
// [MathIsFun]: http://www.mathisfun.com/numbers/addition.html
func Add[N Number](a, b N) N {
	return a + b
}
