package leetcode

import "log"

type Node struct {
	Val  byte
	Cost int
}

func minimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {
	graph := make(map[byte][]Node)
	for i := range original {
		graph[original[i]] = append(graph[original[i]], Node{Val: changed[i], Cost: cost[i]})
	}

	var dfs func(byte, byte, map[byte]struct{}) int
	pathmap := make(map[string]int)
	dfs = func(from, to byte, travelled map[byte]struct{}) int {
		if from == to {
			return 0
		}
		if p, exists := pathmap[string([]byte{from, to})]; exists {
			return p
		}
		m := -1
		if _, exists := travelled[from]; exists {
			return m
		}
		travelled[from] = struct{}{}
		for _, node := range graph[from] {
			t := dfs(node.Val, to, travelled)
			if t == -1 {
				continue
			}
			// pathmap[string([]byte{node.Val, to})] = t
			mpath := node.Cost + t
			if m == -1 {
				m = mpath
			} else {
				m = min(m, mpath)
			}
		}
		delete(travelled, from)
		return m
	}

	res := int64(0)
	travelled := make(map[byte]struct{})
	for i := range source {
		mp := dfs(source[i], target[i], travelled)
		if mp == -1 {
			return -1
		}
		pathmap[string([]byte{source[i], target[i]})] = mp
		res += int64(mp)
	}
	log.Println(pathmap)

	return res
}
