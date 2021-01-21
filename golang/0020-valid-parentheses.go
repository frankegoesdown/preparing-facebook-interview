package main

func main() {

}

func isValid(s string) bool {
	size := len(s)
	stack := make([]byte, size)
	top := 0
	for _, char := range s {
		switch char {
		case '(':
			stack[top] = byte(char + 1)
			top++
		case '[', '{':
			stack[top] = byte(char + 2)
			top++
		case ')', ']', '}':
			if top > 0 && stack[top-1] == byte(char) {
				top--
			} else {
				return false
			}
		}
	}
	return true
}
