package printer

// BinaryTreeInfo 打印接口
type BinaryTreeInfo interface {
	// RootNode who is the root node
	RootNode() any

	// LNode how to get the left child of the node
	LNode() any

	// RNode how to get the right child of the node
	RNode() any

	// ToString how to print the node
	ToString() any

	// ColorOf what's color of the node (true-red、false-black)
	ColorOf() bool
}
