package main

import "fmt"

func main() {
	fmt.Println(myPow(float64(2), 10))
}
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if n < 0 {
		return 1.0 / myPow(x, -n)
	}

	t := myPow(x, n/2)
	fmt.Println(t)
	if n%2 == 1 {
		return t * t * x
	} else {
		return t * t
	}
}
