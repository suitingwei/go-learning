package algorithms

import (
	. "github.com/smartystreets/goconvey/convey"
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
	})
}
