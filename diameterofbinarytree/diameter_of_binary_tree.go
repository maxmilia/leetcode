package diameterofbinarytree

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ans int

func diameterOfBinaryTree(root *TreeNode) int {
	ans = 1
	depth(root)
	return ans - 1
}

func depth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	l := depth(node.Left)
	r := depth(node.Right)
	ans = max(ans, l+r+1)
	return max(l, r) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
