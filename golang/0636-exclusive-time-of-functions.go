package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(exclusiveTime(2, []string{"0:start:0", "1:start:2", "1:end:5", "0:end:6"}))
}

func parseLog(s string) (id int, state string, time int) {
	res := strings.Split(s, ":")
	id, _ = strconv.Atoi(res[0])
	state = res[1]
	time, _ = strconv.Atoi(res[2])
	return
}

func exclusiveTime(n int, logs []string) []int {
	if len(logs) == 0 {
		return make([]int, n, n)
	}

	stack := make([]int, 0, n)
	res := make([]int, n, n)

	// This is used to track the next cycle based on the current log entry.
	lastCoveredTime := 0

	for i := 0; i < len(logs); i++ {
		id, state, time := parseLog(logs[i])
		if state == "start" {
			// Update the top function
			if len(stack) > 0 {
				res[stack[len(stack)-1]] += time - lastCoveredTime
			}
			stack = append(stack, id)
			lastCoveredTime = time // The stack top has covered all time till `time`
		} else {
			res[stack[len(stack)-1]] += time - lastCoveredTime + 1
			stack = stack[:len(stack)-1] // Pop the function as it ended
			lastCoveredTime = time + 1   // The last log entry has covered all time till the end time of the this entry.
		}
		fmt.Println(stack)
		fmt.Println(time)
		fmt.Println(lastCoveredTime)
		fmt.Println("=================")
	}

	return res
}
