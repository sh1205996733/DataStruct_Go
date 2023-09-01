package recursion

import "fmt"

// 汉诺塔
// 实现把 A 的 n 个盘子移动到 C（盘子编号是 [1, n] ）
// 每次只能移动1个盘子且大盘子只能放在小盘子下面

// 将 n 个碟子从 p1 挪动到 p3, p2是中间盘子
func hanio(n int, p1, p2, p3 string) {
	if n == 1 {
		move(n, p1, p3)
		return
	}
	hanio(n-1, p1, p3, p2) // 将 n-1 个碟子从 p1 挪动到 p2, p3是中间盘子
	move(n, p1, p3)
	hanio(n-1, p2, p1, p3) // 将 n-1 个碟子从 p2 挪动到 p3, p1是中间盘子
}

// 将 no 号盘子从 from 移动到 to
func move(n int, from, to string) {
	fmt.Println(fmt.Sprintf("将%d号盘子从%s移动到%s", n, from, to))
}
