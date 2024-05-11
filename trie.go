package leetcode

type TrieNode struct {
	Val      byte
	Children map[byte]*TrieNode
}

func NewTrieTree(s string) *TrieNode {
	root := &TrieNode{Children: make(map[byte]*TrieNode)}
	root.Insert(s)
	return root
}

func (root *TrieNode) Insert(s string) {
	ss := s
	p := root
	for ss != "" {
		if p.Children[ss[0]] == nil {
			p.Children[ss[0]] = &TrieNode{Val: ss[0], Children: make(map[byte]*TrieNode)}
		}
		p = p.Children[ss[0]]
		ss = ss[1:]
	}
}

func (root *TrieNode) Search(s string) bool {
	p := root
	for {
		if s == "" {
			break
		}

		if p.Children[s[0]] == nil {
			return false
		}

		p = p.Children[s[0]]
		s = s[1:]
	}
	return true
}

func (root *TrieNode) SearchPrefix(prefix string) int {
	p := root
	for {
		if prefix == "" {
			break
		}

		if p.Children[prefix[0]] == nil {
			return 0
		}

		p = p.Children[prefix[0]]
		prefix = prefix[1:]
	}

	var bfs func(n *TrieNode) int
	bfs = func(n *TrieNode) int {
		count := 0
		for _, child := range n.Children {
			if child != nil {
				count++
				count += bfs(child)
			}
		}
		return count
	}

	return bfs(p)
}
