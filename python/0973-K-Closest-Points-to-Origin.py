import heapq


class Solution(object):
    def kClosest(self, pairs, k):
        nums = []
        self.getDist(pairs, nums)
        heap = nums[:k]
        heapq.heapify(heap)
        print(nums)
        print(heap)
        for i in range(k, len(nums)):
            if -nums[i][0] < -heap[0][0]:
                heapq.heapreplace(heap, nums[i])

        res = []
        for num, x, y in heap:
            res.append([x, y])
        return res

    def getDist(self, pairs, nums):
        for pair in pairs:
            x, y = pair
            dist = (x ** 2 + y ** 2) ** 0.5
            nums.append((-dist, x, y))


if __name__ == '__main__':
    # begin
    s = Solution()
    print(s.kClosest([[1, 3], [-2, 2]], 1))
