// Task Description: Given an integer n (3 < n < 10^9), find the length of the smallest list of perfect squares which add up to n.
// Come up with the best alogrithm you can; you'll need it.

// Example - sum_of_squares(17) = 2
// 17 = 16 + 1

// 1. Create a fucntion that takes in a number.
// 2. Find out the biggest square that can fit within it.
// 3. Recursion with the remainder and return.

package sumofperfectsquares

import (
	"fmt"
	"strconv"
)

type keyPairValue struct {
	key uint64
	value uint64
}

type solution struct {
	length uint64
	perfectSquares []uint64
}

func SumOfSquares(n uint64) uint64 {
	// Declare an array and store all the possible squares which fit into n.
	var squareValuesWhichFit []keyPairValue; 

	// This contains all the solution for the given lengths.
	var answers []solution;

	var index uint64 = 1;
	for {
		var indexSquared = index * index;
		if n >= uint64(indexSquared) {
			possibleValue := keyPairValue {
				key: index,
				value: indexSquared,
			}
			fmt.Println("Possible value: [" + strconv.Itoa(int(possibleValue.key)) + ", " +  strconv.Itoa(int(possibleValue.value)) + "]");

			squareValuesWhichFit = append(squareValuesWhichFit, possibleValue);
			index++;
			continue;
		} else {
			// Square of the index > n, therefore we have all possible values. 
			break;
		}
	}

	// Based on length of squareValuesWhichFit, we find the minimun per length and then reduce
	var length = len(squareValuesWhichFit) - 1;
	for i := length; i > 0; i-- {		
		var possibleValues []uint64;
		for j := 0; j <= i; j++ {
			possibleValues = append(possibleValues, squareValuesWhichFit[j].value);
		}
		
		// Length solution is an array of squares.
		var lengthSolution = findSolution(n, possibleValues);
		// Grouped values for the value.
		lengthAnswer := solution {
			length: uint64(i),
			perfectSquares: lengthSolution,
		};

		answers = append(answers, lengthAnswer);
	}

	
	// Find the element with the smallest length in answers.
	fmt.Println(answers);
	answer := 0;
	for i, a := range answers {
		l := len(a.perfectSquares);
		if answer > l {
			answer = l;
		} else if  i == 0  {
			answer = l;
		}
	}
	return uint64(answer);
}

func findSolution(n uint64, valuesToPlayWith []uint64) []uint64 {
	// Add some optimisation here.
	if n == 0 {
		return []uint64{};
	}

	solution := []uint64{};
	for i := len(valuesToPlayWith) - 1; i >= 0; i-- {
		var element uint64 = valuesToPlayWith[i];
		if n >= element {
			solution = append(solution, element);
			newValue := n - element;
			result := findSolution(newValue, valuesToPlayWith);
			return append(solution, result...);
		}
	}

	return solution;
}