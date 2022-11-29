package performance

import (
	"bytes"
	"fmt"
	"math/rand"
	"strconv"
	"testing"
)

//BenchmarkFmt-8   	 6088726	       191.1 ns/op
func BenchmarkFmt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%d", rand.Int())
	}
}

//BenchmarkStrconv-8   	 8886268	       128.5 ns/op
func BenchmarkStrconv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strconv.Itoa(rand.Int())
	}
}

func BenchmarkCreateSliceWithConstString(b *testing.B) {
	w := bytes.NewBuffer(nil)
	//每次都初始化slice，即便使用了固定的字符串，但是这个string->[]byte每次都会有开销
	for i := 0; i < b.N; i++ {
		w.Write([]byte("hello world"))
	}
}

func BenchmarkCreateSliceWithOneString(b *testing.B) {
	str := "hello world"
	//这句语法是错误的,因为每一个string，都是不可变的，immutable，因为内部存储的时候，会把string的标量存储在某个内存区域。还会全局缓存
	//str[1] ='b'

	//子字符串，会共享父字符串的内存，也会导致如果子字符串不释放一直占用，那么父字符串也无法回收内存
	//fmt.Println(str[2:3])

	//每次都初始化slice，即便使用了固定的字符串，但是这个string->[]byte每次都会有开销
	w := bytes.NewBuffer(nil)
	for i := 0; i < b.N; i++ {
		w.Write([]byte(str))
	}
}
