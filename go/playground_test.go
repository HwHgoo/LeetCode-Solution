package leetcode_test

import (
	"log"
	"math"
	"strings"
	"testing"
)

func TestInf(t *testing.T) {
	inf := int(math.Inf(1))
	log.Println(math.IsInf(float64(inf), 1))
}

func TestSortSearch(t *testing.T) {
}

func TestStringBuilder(t *testing.T) {
	b := strings.Builder{}
	b.Reset()
}

func TestGoDs(t *testing.T) {
}
