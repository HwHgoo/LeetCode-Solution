package leetcode

import (
	"testing"

	"github.com/kr/pretty"
)

func TestTrie(t *testing.T) {
	root := NewTrieTree("abc")
	root.Insert("cba")
	root.Insert("cjb")
	pretty.Logln(root.SearchPrefix("c"))
}
