package leetcode

/*
 * @lc app=leetcode.cn id=71 lang=golang
 *
 * [71] Simplify Path
 */

// @lc code=start
func simplifyPath(path string) string {
	st := make([]byte, len(path))
	st[0] = '/'
	r, n := 1, len(path)
	top := 1
	for r < n {
		if path[r] == '/' {
			r++
		} else if path[r] == '.' && (r+1 == n || path[r+1] == '/') {
			r++
		} else if path[r] == '.' && path[r+1] == '.' && (r+2 == n || path[r+2] == '/') {
			r += 2
			if top > 1 {
				top--
				for top > 1 && st[top] != '/' {
					top--
				}
			}
		} else {
			if top != 1 {
				st[top] = '/'
				top++
			}
			for ; r < n && path[r] != '/'; r++ {
				st[top] = path[r]
				top++
			}
		}
	}
	return string(st[:top])
}

// @lc code=end

/*

	find_parent_root := func(top int) int {
		if top-3 == 0 {
			return 0
		}

		top -= 4
		for top > 0 && st[top] != '/' {
			top--
		}
		return top
	}
	for _, c := range path[1:] {
		if c == '/' {
			if st[top-1] == '/' {
				continue
			} else if top-2 >= 0 && string(st[top-2:top]) == "/." {
				top--
				continue
			} else if top-3 >= 0 && string(st[top-3:top]) == "/.." {
				r := find_parent_root(top)
				top = r + 1
				continue
			}
		}
		st[top] = byte(c)
		top++
	}

	if st[top-1] == '/' && top != 1 {
		top--
	}
	return string(st[:top])
*/
