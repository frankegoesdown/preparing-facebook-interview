package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(alienOrder([]string{"wrt", "wrf", "er", "ett", "rftt"}))
}

func alienOrder(words []string) string {
	m := map[byte][]byte{}

	for _, word := range words {
		for _, r := range word {
			m[byte(r)] = []byte{}
		}
	}
	fmt.Println(m)
	for i := 0; i+1 < len(words); i++ {
		word1, word2 := words[i], words[i+1]

		if len(word1) > len(word2) && strings.HasPrefix(word1, word2) {
			return ""
		}

		for j := 0; j < min(len(word1), len(word2)); j++ {
			if word1[j] != word2[j] {
				m[word2[j]] = append(m[word2[j]], word1[j])
				break
			}
		}
	}
	fmt.Println(m)
	seen := map[byte]bool{}

	output := ""

	for r, _ := range m {
		result := dfs(r, seen, m, &output)
		if !result {
			return ""
		}
	}

	if len(output) > len(m) {
		return ""
	}

	return output
}

func dfs(r byte, seen map[byte]bool, m map[byte][]byte, output *string) bool {
	if v, ok := seen[r]; ok {
		return v
	}

	for _, sr := range m[r] {
		result := dfs(sr, seen, m, output)
		if !result {
			return false
		}
	}

	seen[r] = true
	*output += string(r)

	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
