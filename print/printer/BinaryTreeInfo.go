package printer

// 打印接口
type BinaryTreeInfo interface {
	/**
	 * who is the root node
	 */
	RootNode() interface{}
	/**
	 * how to get the left child of the node
	 */
	LNode() interface{}
	/**
	 * how to get the right child of the node
	 */
	RNode() interface{}
	/**
	 * how to print the node
	 */
	ToString() interface{}
	/**
	 * what's color of the node (true-red、false-black)
	 */
	ColorOf() bool
}
