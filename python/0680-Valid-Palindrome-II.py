class Solution(object):
    def validPalindrome(self, s):
        def is_pali_range(i, j):
            print(i)
            print(j)
            return all(s[k] == s[j - k + i] for k in range(i, j))

        for i in range(len(s) // 2):
            if s[i] != s[~i]:
                j = len(s) - 1 - i
                return is_pali_range(i + 1, j) or is_pali_range(i, j - 1)
        return True


if __name__ == "__main__":
    s = Solution()
    print(s.validPalindrome("abca"))
