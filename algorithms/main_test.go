package algorithms

import (
	. "github.com/smartystreets/goconvey/convey"
	"math"
	"testing"
)

func add(x, y int) int {
	result := x + y
	//fmt.Println("Result:",result)
	return result
}

func TestAdd(t *testing.T) {

	Convey("测试两数字相加", t, func() {
		So(add(1, 2), ShouldEqual, 3)
		So(add(1, 4), ShouldEqual, 5)
	})
	Convey("测试负数相加", t, func() {
		So(add(1, -2), ShouldEqual, -1)
		So(add(1, -4), ShouldEqual, -3)
	})
}

func TestAdd2(t *testing.T) {

	Convey("测试两数字相加", t, func() {
		So(add(1, 2), ShouldEqual, 3)
		So(add(1, 4), ShouldEqual, 5)
	})
	Convey("测试负数相加", t, func() {
		So(add(1, -2), ShouldEqual, -1)
		So(add(1, -4), ShouldEqual, -3)
	})

	Convey("测试相加溢出相加", t, func() {
		So(add(math.MaxInt64, 2), ShouldBeLessThan, 0)
	})
}
