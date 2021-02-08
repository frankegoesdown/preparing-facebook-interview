class Solution:
    def subarraySum(self, nums, k):
        preSums = {0: 1}
        s = 0
        res = 0
        for num in nums:
            s += num
            res += preSums.get(s - k, 0)
            preSums[s] = preSums.get(s, 0) + 1
            print(s)
            print(res)
            print(preSums)
        return res


if __name__ == '__main__':
    # begin
    s = Solution()
    print(s.subarraySum([1, 1, 1], 2))
