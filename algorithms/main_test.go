package main

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

func TestLargeGroupPositions(t *testing.T) {

	Convey("多个满足条件", t, func() {

		Convey("连续出现", func() {
			So(largeGroupPositions("abcdddeeeeaabbbcd"), ShouldResemble, [][]int{{3, 5}, {6, 9}, {12, 14}})
		})

		Convey("不连续出现", func() {
			So(largeGroupPositions("abccceaaabbcd"), ShouldResemble, [][]int{{2, 4}, {6, 8}})
		})
	})
	Convey("只有一个满足条件", t, func() {

		Convey("长度大于 3", func() {
			So(largeGroupPositions("abbxxxxzzy"), ShouldResemble, [][]int{{3, 6}})
			So(largeGroupPositions("abbxxx"), ShouldResemble, [][]int{{3, 5}})
			So(largeGroupPositions("xxxx"), ShouldResemble, [][]int{{0, 3}})
		})

		Convey("长度等于 3", func() {
			So(largeGroupPositions("aaa"), ShouldResemble, [][]int{{0, 2}})
		})

	})

	Convey("不满足条件", t, func() {

		Convey("长度小于3", func() {
			Convey("空字符串", func() {
				So(largeGroupPositions(""), ShouldResemble, [][]int{})
			})
			Convey("非空字符串", func() {
				So(largeGroupPositions("abc"), ShouldResemble, [][]int{})
			})
		})

		Convey("长度大于3", func() {
			So(largeGroupPositions("abcdhaoniqwe"), ShouldResemble, [][]int{})
		})

	})

}
