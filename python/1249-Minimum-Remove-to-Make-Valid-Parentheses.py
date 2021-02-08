class Solution:
    def minRemoveToMakeValid(self, s: str) -> str:
        res = []
        first = []
        l = 0
        total = 0
        for c in s:
            if c == '(':
                l += 1
                total += 1
            if c == ')':
                if l == 0:
                    continue
                l -= 1
                print(l)
            first.append(c)
        print(first)
        removed = total - l
        print(removed)
        for c in first:
            if c == '(':
                removed -= 1
                if removed < 0:
                    continue
            res.append(c)
        return "".join(res)


if __name__ == '__main__':
    # begin
    s = Solution()
    print(s.minRemoveToMakeValid("lee(t(c)o)de)"))
