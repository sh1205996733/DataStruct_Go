package kmp

// 蛮力算法优化（此前实现的蛮力算法，在恰当的时候可以提前退出，减少比较次数）
func indexOf2(text, pattern string) int {
	if text == "" || pattern == "" {
		return -1
	}
	tlen := len(text)
	plen := len(pattern)
	if tlen == 0 || plen == 0 || tlen < plen {
		return -1
	}
	tmax := tlen - plen
	for ti := 0; ti < tmax; ti++ {
		pi := 0
		for ; pi < plen; pi++ {
			if text[ti+pi] != pattern[pi] {
				break
			}
		}
		if pi == plen {
			return ti
		}
	}
	return -1
}

// 蛮力算法优化（此前实现的蛮力算法，在恰当的时候可以提前退出，减少比较次数）
func indexOf1(text, pattern string) int {
	if text == "" || pattern == "" {
		return -1
	}
	tlen := len(text)
	plen := len(pattern)
	if tlen == 0 || plen == 0 || tlen < plen {
		return -1
	}
	ti, pi, tmax := 0, 0, tlen-plen
	// 因此，ti 的退出条件可以从 ti < tlen 改为 ti – pi <= tlen – plen
	// ti – pi 是指每一轮比较中 text 首个比较字符的位置
	for ti-pi <= tmax && pi < plen {
		if text[ti+pi] == pattern[pi] {
			pi++
		} else {
			ti++
			pi = 0
		}
	}
	if pi == plen {
		return ti - pi
	}
	return -1
}

// 蛮力算法 以字符为单位，从左到右移动模式串，直到匹配成功
func indexOf(text, pattern string) int {
	if text == "" || pattern == "" {
		return -1
	}
	tlen := len(text)
	plen := len(pattern)
	if tlen == 0 || plen == 0 || tlen < plen {
		return -1
	}
	ti, pi := 0, 0
	for pi < plen && ti < tlen {
		if text[ti] == pattern[pi] {
			ti++
			pi++
		} else {
			ti -= pi - 1
			pi = 0
		}
	}
	if pi == plen {
		return ti - pi
	}
	return -1
}
