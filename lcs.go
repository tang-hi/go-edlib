package edlib

import "errors"

// LCS takes two strings and compute their LCS(Longuest Subsequence Problem)
func LCS(str1, str2 string) int {
	// Convert strings to rune array to handle no-ASCII characters
	runeStr1 := []rune(str1)
	runeStr2 := []rune(str2)

	if len(runeStr1) == 0 || len(runeStr2) == 0 {
		return 0
	} else if equal(runeStr1, runeStr2) {
		return len(runeStr1)
	}

	lcsMatrix := lcsProcess(runeStr1, runeStr2)
	return lcsMatrix[len(runeStr1)][len(runeStr2)]
}

func lcsProcess(runeStr1, runeStr2 []rune) [][]int {
	// 2D Array that will contain str1 and str2 LCS
	lcsMatrix := make([][]int, len(runeStr1)+1)
	for i := 0; i <= len(runeStr1); i++ {
		lcsMatrix[i] = make([]int, len(runeStr2)+1)
		for j := 0; j <= len(runeStr2); j++ {
			lcsMatrix[i][j] = 0
		}
	}

	for i := 1; i <= len(runeStr1); i++ {
		for j := 1; j <= len(runeStr2); j++ {
			if runeStr1[i-1] == runeStr2[j-1] {
				lcsMatrix[i][j] = lcsMatrix[i-1][j-1] + 1
			} else {
				lcsMatrix[i][j] = max(lcsMatrix[i][j-1], lcsMatrix[i-1][j])
			}
		}
	}

	return lcsMatrix
}

// LCSBacktrack returns all choices taken during LCS process
func LCSBacktrack(str1, str2 string) (string, error) {
	runeStr1 := []rune(str1)
	runeStr2 := []rune(str2)

	if len(runeStr1) == 0 || len(runeStr2) == 0 {
		return "", errors.New("Can't process and backtrack any LCS with empty string")
	} else if equal(runeStr1, runeStr2) {
		return str1, nil
	}

	lcsMatrix := make([][]int, len(runeStr1)+1)
	for i := 0; i <= len(runeStr1); i++ {
		lcsMatrix[i] = make([]int, len(runeStr2)+1)
		for j := 0; j <= len(runeStr2); j++ {
			lcsMatrix[i][j] = 0
		}
	}

	return processLCSBacktrack(str1, str2, lcsProcess(runeStr1, runeStr2), len(str1), len(str2)), nil
}

func processLCSBacktrack(str1 string, str2 string, lcsMatrix [][]int, m, n int) string {
	// Convert strings to rune array to handle no-ASCII characters
	runeStr1 := []rune(str1)
	runeStr2 := []rune(str2)

	if m == 0 || n == 0 {
		return ""
	}
	if runeStr1[m-1] == runeStr2[n-1] {
		return processLCSBacktrack(str1, str2, lcsMatrix, m-1, n-1) + string(str1[m-1])
	}
	if lcsMatrix[m][n-1] > lcsMatrix[m-1][n] {
		return processLCSBacktrack(str1, str2, lcsMatrix, m, n-1)
	}

	return processLCSBacktrack(str1, str2, lcsMatrix, m-1, n)
}

// LCSEditDistance determines the edit distance between two strings using LCS function
// (allow only insert and delete operations)
func LCSEditDistance(str1, str2 string) int {
	if len(str1) == 0 {
		return len(str2)
	} else if len(str2) == 0 {
		return len(str1)
	} else if str1 == str2 {
		return 0
	}

	lcs := LCS(str1, str2)
	return (len(str1) - lcs) + (len(str2) - lcs)
}
