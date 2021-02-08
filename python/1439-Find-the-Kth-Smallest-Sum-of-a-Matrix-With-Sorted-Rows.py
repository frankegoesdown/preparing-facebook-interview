import heapq


def kthSmallest(mat, k):
    """
    :type mat: List[List[int]]
    :type k: int
    :rtype: int
    """
    m, n = len(mat), len(mat[0])
    smallest = sum(mat[i][0] for i in range(m))
    seen = set()
    hp = [(smallest, [0] * m)]
    k -= 1
    print(smallest)
    print(hp)
    while hp:
        total, idx = heapq.heappop(hp)
        if not k:
            return total

        k -= 1
        for i in range(m):
            col = idx[i]
            if col == n - 1:
                continue

            c_total, c_idx = total, idx[:]
            c_total -= mat[i][col]
            c_total += mat[i][col + 1]
            c_idx[i] = col + 1

            if tuple(c_idx) in seen:
                continue
            heapq.heappush(hp, (c_total, c_idx))
            seen.add(tuple(c_idx))

    return -1


print(kthSmallest([[1, 3, 11], [2, 4, 6]], 5))
