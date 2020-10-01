package edlib

import (
	"fmt"
	"math"
	"strings"

	"github.com/hbollon/go-edlib/internal/utils"
)

// CosineSimilarity use cosine algorithm to return a similarity index between string vectors
// Takes two strings as parameters and return an index.
func CosineSimilarity(str1, str2 string) float32 {
	// Split string before rune conversion for cosine calculation
	splittedStr1 := strings.Split(str1, " ")
	splittedStr2 := strings.Split(str2, " ")
	fmt.Printf("%v , %v\n", splittedStr1, splittedStr2)

	// Conversion of plitted string into rune array
	runeStr1 := make([][]rune, len(splittedStr1))
	for i, str := range splittedStr1 {
		runeStr1[i] = []rune(str)
	}
	runeStr2 := make([][]rune, len(splittedStr2))
	for i, str := range splittedStr2 {
		runeStr2[i] = []rune(str)
	}

	var l1, l2 []int
	// Create union keywords slice between input strings
	unionStr := union(splittedStr1, splittedStr2)
	for _, word := range unionStr {
		_, fw := find(runeStr1, word)
		if fw {
			l1 = append(l1, 1)
		} else {
			l1 = append(l1, 0)
		}

		_, fw = find(runeStr2, word)
		if fw {
			l2 = append(l2, 1)
		} else {
			l2 = append(l2, 0)
		}
	}
	fmt.Printf("%v\n%v\n%v\n", unionStr, l1, l2)

	// Compute cosine algorithm
	var cosineSim float32
	for i := 0; i < len(unionStr); i++ {
		cosineSim += float32(l1[i] * l2[i])
	}

	cosineSim = cosineSim / float32(math.Sqrt(float64(sum(l1)*sum(l2))))
	fmt.Printf("Cosine sim between %s/%s : %f\n", str1, str2, cosineSim)
	return cosineSim
}

// Compute union between two string slices, convert result to rune matrix and return it
func union(a, b []string) [][]rune {
	m := make(map[string]bool)
	for _, item := range a {
		m[item] = true
	}
	for _, item := range b {
		if _, ok := m[item]; !ok {
			a = append(a, item)
		}
	}

	// Convert a to rune matrix (with x -> words and y -> characters)
	out := make([][]rune, len(a))
	for i, word := range a {
		out[i] = []rune(word)
	}
	return out
}

// Find takes a rune slice and looks for an element in it. If found it will
// return it's key, otherwise it will return -1 and a bool of false.
func find(slice [][]rune, val []rune) (int, bool) {
	for i, item := range slice {
		if utils.Equal(item, val) {
			return i, true
		}
	}
	return -1, false
}

// Return the elements sum from int slice
func sum(arr []int) int {
	var res int
	for _, v := range arr {
		res += v
	}
	return res
}

func dot(vect1, vect2 []int) float32 {
	var out float32
	for i := 0; i < len(vect1); i++ {
		out += float32(vect1[i] * vect2[i])
	}
	return out
}
