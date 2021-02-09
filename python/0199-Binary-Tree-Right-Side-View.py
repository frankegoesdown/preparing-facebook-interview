from typing import List


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def rightSideView(self, root: TreeNode) -> List[int]:
        if root is None:
            return []
        result = []
        self.dfs(root, result, 0)
        return result

    def dfs(self, node: TreeNode, result: List[int], depth: int) -> None:
        if depth == len(result):
            result.append(node.val)
        for ch in [node.right, node.left]:
            if ch:
                self.dfs(ch, result, depth + 1)


if __name__ == "__main__":
    s = Solution()
    root = TreeNode(1)
    root.left = TreeNode(2)
    root.right = TreeNode(3)
    root.left.left = TreeNode(4)
    root.left.right = TreeNode(5)
    root.right.left = TreeNode(6)
    root.right.right = TreeNode(7)
    root.right.left.right = TreeNode(8)

    print(s.rightSideView(root))
