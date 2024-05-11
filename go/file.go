package leetcode

import (
	"bufio"
	"fmt"
	"os"
	"unsafe"
)

type caseType interface {
	int | []int | string | []string
}

type TestcaseReader[V caseType] struct {
	file   *os.File
	reader *bufio.Reader
}

func open_testcase[V caseType](no int, withTarget bool) *TestcaseReader[V] {
	path := fmt.Sprintf("/home/huwenhao/poj/leetcode/testcases/%d.testcase", no)
	f, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil
	}
	reader := bufio.NewReader(f)
	return &TestcaseReader[V]{
		file:   f,
		reader: reader,
	}
}

func (tr *TestcaseReader[V]) next() V {
	l, _, _ := tr.reader.ReadLine()
	ptr := unsafe.Pointer(&l)
	v := (*V)(ptr)
	return *v
}

func (tr *TestcaseReader[V]) close() {
	tr.file.Close()
}
