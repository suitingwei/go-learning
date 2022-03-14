package types

import (
	"math"
	"testing"
)

//浮点数加法(考虑到浮点数的表示形式，其加减法会有很多不符合常理的场景)
func TestFloatAddition(t *testing.T) {
	a := 1e20
	b := -1e20

	t.Log((a + b) + 3.14)
	t.Log((a + 3.14) + b)
	t.Log(a + 3.14)

}

//整形计算，x^2 >0 一定成立吗？
func TestIntMultipe(t *testing.T) {
	var a int = 4000000
	var b int = 5000000
	t.Log(int32(a * a))
	//整形溢出异常，变得不再是正数
	t.Log(int32(b * b))
	//有符号的in8，二进制标识为：0000 0000
	//其中最高位是符号位，0是负数，1是正数，所以他的表示范围是:
	//[0111,1111 - 1111,1111]
	var int8OverFlow = math.MaxInt8 + 1

	t.Logf("Int8 Max=%d, binary format=%08b", math.MaxInt8, math.MaxInt8)
	t.Logf("Int8 Max Overflow=%d, binary format=%08b", int8(int8OverFlow), int8(int8OverFlow))
	t.Logf("Uint8 Max Overflow=%d, binary format=%08b", uint8(int8OverFlow), uint8(int8OverFlow))
}
