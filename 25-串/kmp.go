package kmp

// 对比蛮力算法，KMP的精妙之处：充分利用了此前比较过的内容，可以很聪明地跳过一些不必要的比较位置
func kmp(text, pattern string) int {
	if text == "" || pattern == "" {
		return -1
	}
	tlen := len(text)
	plen := len(pattern)
	if tlen == 0 || plen == 0 || tlen < plen {
		return -1
	}
	ti, pi, tmax := 0, 0, tlen-plen
	//在暴力算法indexOf之间增加一个next表
	next := next(pattern)
	for ti-pi <= tmax && pi < plen {
		if pi < 0 || text[ti] == pattern[pi] {
			ti++
			pi++
		} else {
			pi = next[pi]
		}
	}
	if pi == plen {
		return ti - pi
	}
	return -1
}

// KMP – next表的优化
func next(pattern string) []int {
	plen := len(pattern)
	next := make([]int, plen)
	i, imax := 0, plen-1
	next[i] = -1
	n := next[i]
	for i < imax {
		if n < 0 || pattern[i] == pattern[n] {
			i++
			n++
			if pattern[i] == pattern[n] {
				next[i] = next[n]
			} else {
				next[i] = n
			}
		} else {
			n = next[n]
		}
	}
	return next
}

func next0(pattern string) []int {
	plen := len(pattern)
	next := make([]int, plen)
	i, imax := 0, plen-1
	next[i] = -1
	n := next[i]
	for i < imax {
		if n < 0 || pattern[i] == pattern[n] {
			i++
			n++
			next[i] = n
		} else {
			n = next[n]
		}
	}
	return next
}
