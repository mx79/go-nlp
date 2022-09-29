package distance

import (
	"github.com/mx79/go-nlp/clean"
	"math"
)

// checkEntries is a utility function to check
// strings before calculating Levenshtein scores
func checkEntries(a, b string) int {
	if len(a) == 0 {
		return len(b)
	}
	if len(b) == 0 {
		return len(a)
	}
	if a == b {
		return 0
	}
	return -1
}

// min
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//// Levenshtein distance between words at the character level
//func Levenshtein(a, b string) int {
//	// Checking if it is necessary or not to calculate the Levenshtein
//	if check := checkEntries(a, b); check != -1 {
//		return check
//	}
//	// D is an array of lengthString1+1 rows and lengthString2+1 columns
//	// D is indexed from 0, strings from 1
//	// i and j are indexes
//	// cost is the substitution cost of the algorithm
//	D := make([][]int, len(a)+1)
//	cost := make([][]int, len(a))
//	var i, j int
//	// Init D matrix
//	for idx := range D {
//		D[idx] = make([]int, len(b)+1)
//	}
//	// Init cost matrix
//	for idx := range cost {
//		cost[idx] = make([]int, len(b))
//	}
//	// Charging matrix D with our strings a and b
//	for i = 0; i < len(D); i++ {
//		D[i][0] = i
//	}
//	for j = 0; j < len(D[0]); j++ {
//		D[0][j] = j
//	}
//	// Charging matrix cost with our strings a and b
//	for i = 0; i < len(a); i++ {
//		for j = 0; j < len(b); j++ {
//			if a[i] == b[j] {
//				cost[i][j] = 0
//			} else {
//				cost[i][j] = 1
//			}
//		}
//	}
//	fmt.Printf("cost matrix %v\n", cost)
//	// Levenshtein algorithm
//	for i = 1; i < len(a); i++ {
//		for j = 1; j < len(b); j++ {
//			if a[i] == b[j] {
//				// cost[i][j] = 0
//			} else {
//				// cost[i][j] = 1
//				D[i][j] = min(min(D[i-1][j]+1, D[i][j-1]+1), D[i-1][j-1]+cost[i-1][j-1])
//			}
//		}
//	}
//	fmt.Printf("D matrix %v\n", D)
//	return D[len(a)][len(b)]
//}

// WordErrorRate is the  distance between
// two sentences at the word level
func WordErrorRate(a, b string) float64 {
	var (
		res int
		str []string
	)
	if check := checkEntries(a, b); check != -1 {
		return float64(check)
	}
	elemA := clean.Tokenize(a, true)
	elemB := clean.Tokenize(b, true)
	if len(elemA) < len(elemB) {
		str = elemA
		res += len(elemB) - len(elemA)
	} else {
		str = elemB
		res += len(elemB) - len(elemA)
	}
	for index := range str {
		if elemA[index] != elemB[index] {
			res++
		}
	}
	return math.Abs(float64(res)) / float64(len(str))
}
