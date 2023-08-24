package printer

const (
	LEVEL_ORDER = "LEVEL_ORDER"
	INORDER     = "INORDER"
)

func Print(info BinaryTreeInfo, style ...string) {
	if info == nil || info.RootNode() == nil {
		return
	}
	printer(info, style...).Print()
}

func Println(info BinaryTreeInfo, style ...string) {
	if info == nil || info.RootNode() == nil {
		return
	}
	printer(info, style...).Println()
}

func printer(info BinaryTreeInfo, style ...string) basePrinter {
	if len(style) == 1 {
		if style[0] == INORDER {
			return initInorderPrinter(info)
		} else if style[0] == LEVEL_ORDER {
			return initLevelOrderPrinter(info)
		} else {
			return initLevelOrderPrinter(info)
		}
	} else {
		return initLevelOrderPrinter(info)
	}
}
