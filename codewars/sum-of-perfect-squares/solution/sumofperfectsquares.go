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
	"math"
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
	if isDirectSquareNumber(n) {
		return 1;
	}

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
			// fmt.Println("Possible value: [" + strconv.Itoa(int(possibleValue.key)) + ", " +  strconv.Itoa(int(possibleValue.value)) + "]");

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
		fmt.Println(i);

		var possibleValues []uint64;
		for j := 0; j <= i; j++ {
			possibleValues = append(possibleValues, squareValuesWhichFit[j].value);
		}
		
		// Length solution is an array of squares.
		var lengthSolution = findSolution(n, possibleValues, answers);
		// Grouped values for the value.
		lengthAnswer := solution {
			length: uint64(i),
			perfectSquares: lengthSolution,
		};

		answers = append(answers, lengthAnswer);
	}

	return findShortestAnswer(answers);
}

func isDirectSquareNumber(n uint64) bool {
	p := math.Sqrt(float64(n));
	return  p == float64(int(p));
}

func findSolution(n uint64, valuesToPlayWith []uint64, answers []solution) []uint64 {
	// Add some optimisation here.
	if n == 0 {
		return []uint64{};
	}

	solution := []uint64{};
	lengthOfValuesToPlayWith := len(valuesToPlayWith);
	for i := lengthOfValuesToPlayWith - 1; i >= 0; i-- {
		var element uint64 = valuesToPlayWith[i];
		if n >= element {
			solution = append(solution, element);
			newValue := n - element;
			// Idea: here if we have 661878529 and then left with 37174 all the square inbetween are useless.
			// 1. Based on the new value we want to exclude of the values in the possible slice which are greater.
			// 2. This should dramatically reduce the length of the slice passed down into the recursive func's.
			
			highArrayIndex := findPossibleValueIndex(newValue, valuesToPlayWith);
			newValuesToPlayWith := valuesToPlayWith[0 : highArrayIndex];
			result := findSolution(newValue, newValuesToPlayWith, answers);
			solution = append(solution, result...);

			// 3. Use the answers collection that we are building up, this can allow us to discard answers quickly.
			// currentShortestAnswer := findShortestAnswer(answers);
			// if uint64(len(solution)) >= currentShortestAnswer {
			// 	// discard.
			// 	return solution;
			// }

			return solution;
		}
	}

	return solution;
}

func findPossibleValueIndex(newValue uint64, valuesToPlayWith []uint64) int {
	for i, value := range valuesToPlayWith {
		if value >= newValue {
			return i;
		}
	}

	return len(valuesToPlayWith);
}

func findShortestAnswer(answers []solution) uint64 {
	// Find the element with the smallest length in answers.
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