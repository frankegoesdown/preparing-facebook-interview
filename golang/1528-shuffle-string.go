package main

func main() {
	
}

func restoreString(s string, indices []int) string {
	resultBytes := make([]byte,len(s))

	for i, index := range indices {
		resultBytes[index] = s[i]
	}

	return string(resultBytes)
}
