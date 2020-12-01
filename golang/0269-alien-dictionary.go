package main

import "strings"

func main() {

}

func alienOrder(words []string) string {
	m := map[rune][]rune{}

	for _, word := range words {
		for _, r := range word {
			m[r] = []rune{}
		}
	}

	for i := 0; i+1 < len(words); i++ {
		word1, word2 := words[i], words[i+1]

		if len(word1) > len(word2) && strings.HasPrefix(word1, word2) {
			return ""
		}

		for j := 0; j < min(len(word1), len(word2)); j++ {
			if word1[j] != word2[j] {
				m[rune(word2[j])] = append(m[rune(word2[j])], rune(word1[j]))
				break
			}
		}
	}

	seen := map[rune]bool{}

	var dfs func(r rune) bool

	output := ""

	dfs = func(r rune) bool {
		if v, ok := seen[r]; ok {
			return v
		}

		seen[r] = false

		for _, sr := range m[r] {
			result := dfs(sr)
			if !result {
				return false
			}
		}

		seen[r] = true
		output += string(r)

		return true
	}

	for r, _ := range m {
		result := dfs(r)
		if !result {
			return ""
		}
	}

	if len(output) > len(m) {
		return ""
	}

	return output
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
