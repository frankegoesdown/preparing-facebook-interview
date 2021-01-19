package main

func main() {


}

func wordBreak(s string, w []string) []string {
	m := make(map[string]bool, len(w))
	for _, v := range w {
		m[v] = true
	}
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := i+1; j < len(s); j++ {
			if m[s[i:j]] {
				for k := j; k < len(s); k++ {
					if dp[j][k] {
						dp[i][j-1] = true
						break
					}
				}
			}
		}
		if m[s[i:]] {
			dp[i][len(s) - 1] = true
		}
	}
	var out []string
	dfs(s, dp, 0, &out, &bytes.Buffer{})
	return out
}

func dfs(s string, dp[][]bool, ind int, out *[]string, buf *bytes.Buffer) {
	if ind == len(s) {
		*out = append(*out, buf.String())
		return
	}
	for i := ind; i < len(s); i++ {
		if dp[ind][i] {
			if buf.Len() > 0 {
				buf.WriteString(" ")
			}
			buf.WriteString(s[ind:i+1])
			dfs(s, dp, i + 1, out, buf)
			buf.Truncate(buf.Len() - (i+1-ind))
			if buf.Len() > 0 {
				buf.Truncate(buf.Len() - 1)
			}
		}
	}
}

func wordBreak2(s string, wordDict []string) []string {
	mp := make(map[string][]string)
	return dfs2(s, wordDict, mp)
}

func dfs2(s string, wordDict []string, mp map[string][]string) []string {
	_, ok := mp[s]
	if ok {
		return mp[s]
	}

	res := make([]string, 0)
	if len(s) == 0 {
		res = append(res, "")
		return res
	}
	for _,word:= range wordDict {
		if strings.HasPrefix(s, word) {
			sublist := dfs2(s[len(word):], wordDict, mp)
			for _, sub:= range sublist {
				a := " "
				if len(sub) == 0 {
					a = ""
				}
				res = append(res, word + a + sub)
			}
		}
	}
	mp[s] = res
	return res
}


