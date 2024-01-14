package tree_node

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}

func (n *TreeNode) Insert(val int) *TreeNode {
	if n == nil {
		return NewTreeNode(val)
	}

	if val < n.Val {
		n.Right = n.Right.Insert(val)
	} else {
		n.Left = n.Left.Insert(val)
	}

	return n
}

func (n *TreeNode) ToList() []int {
	if n == nil {
		return []int{}
	}

	return append([]int{n.Val}, append(n.Left.ToList(), n.Right.ToList()...)...)
}

var NULL = -1 << 63

func Ints2TreeNode(ints []int) *TreeNode {
	n := len(ints)
	if n == 0 {
		return nil
	}

	root := &TreeNode{
		Val: ints[0],
	}

	queue := make([]*TreeNode, 1, n*2)
	queue[0] = root

	i := 1
	for i < n {
		node := queue[0]
		queue = queue[1:]

		if i < n && ints[i] != NULL {
			node.Left = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < n && ints[i] != NULL {
			node.Right = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}
