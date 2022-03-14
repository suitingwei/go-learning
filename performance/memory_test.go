package main

import (
	"testing"
	"unsafe"
)

func TestEmptyStruct(t *testing.T) {
	t.Logf("empty struct size=%d", unsafe.Sizeof(struct{}{}))
	t.Logf("empty string slice size=%d", unsafe.Sizeof([]string{}))
	t.Logf("empty int slice size=%d", unsafe.Sizeof([]int{}))
	t.Logf("empty string map size=%d", unsafe.Sizeof(map[string]string{}))
	t.Logf("empty int map size=%d", unsafe.Sizeof(map[int]int{}))

	//pointers
	intPtr := new(int)
	strPtr := new(string)
	t.Logf("int pointer size=%d", unsafe.Sizeof(intPtr))
	t.Logf("string pointer size=%d", unsafe.Sizeof(strPtr))
}
