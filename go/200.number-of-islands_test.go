package leetcode

import (
	"log"
	"testing"
)

func TestNumIsLands(t *testing.T) {
	grid1 := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}

	grid2 := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}

	grid3 := [][]byte{
		{'1', '0', '1', '1', '1'},
		{'1', '0', '1', '0', '1'},
		{'1', '1', '1', '0', '1'},
	}

	log.Println("grid1 islands: ", numIslands(grid1))
	log.Println("grid2 islands: ", numIslands(grid2))
	log.Println("grid3 islands: ", numIslands(grid3))
}
