class Solution:
    def isAlienSorted(self, words, order):
        """
        :type words: List[str]
        :type order: str
        :rtype: bool
        """
        dic = {}
        for i, c in enumerate(order):
            dic[c] = i
        print(dic)
        for i in range(len(words)-1):
            cur = words[i]
            nex = words[i+1]
            print(cur)
            print(nex)
            print("----------")
            for j in range(min(len(cur), len(nex))):
                print(cur[j])
                print(nex[j])
                print("======")
                if cur[j] != nex[j]:
                    if dic[cur[j]] > dic[nex[j]]:
                        return False
                    break
            else:
                if len(cur) > len(nex):
                    return False
        return True


if __name__ == '__main__':
    # begin
    s = Solution()
    print(s.isAlienSorted(["hello", "leetcode"], "hlabcdefgijkmnopqrstuvwxyz"))
