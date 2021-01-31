package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(groupStrings2([]string{"abc", "bcd", "acef", "xyz", "az", "ba", "a", "z"}))
}

func groupStrings(strings []string) [][]string {
	m := make(map[string][]string)
	for _, s := range strings {
		fmt.Println(s)
		ss := shift(s, s[0]-'a')
		fmt.Println("------")
		m[ss] = append(m[ss], s)
	}
	var res [][]string
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

func shift(s string, count byte) string {
	var sb bytes.Buffer
	for i := 0; i < len(s); i++ {
		ch := (s[i]-'a'+26-count)%26 + 'a'
		fmt.Println(ch)
		sb.WriteByte(ch)
	}
	return sb.String()
}

func groupStrings2(strings []string) [][]string {
	m := make(map[string][]string)
	result := make([][]string, 0)

	for i := 0; i < len(strings); i++ {
		cur := strings[i]
		diff := cur[0] - 'a'
		byteArr := make([]byte, 0)
		for k := 0; k < len(cur); k++ {
			fmt.Println(cur)
			fmt.Println(k)
			fmt.Println(diff)
			if cur[k]-diff < 'a' {
				fmt.Println(cur[k]-diff+26)
				byteArr = append(byteArr, cur[k]-diff+26)
			} else {
				byteArr = append(byteArr, cur[k]-diff)
			}
		}
		fmt.Println("-----")
		fmt.Println(byteArr)
		fmt.Println("========")
		m[string(byteArr)] = append(m[string(byteArr)], cur)
	}

	for _, val := range m {
		fmt.Println(val)
		result = append(result, val)
	}
	return result
}
