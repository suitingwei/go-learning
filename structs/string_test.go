package structs

import "testing"

func testStringCopy() string {
	var s string

	for i := 0; i < 1000; i++ {
		s += "a"
	}
	return s
}

func BenchmarkTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testStringCopy()
	}
}

func TestSlice(t *testing.T) {

	data := [3]int{1, 2, 3}

	t.Log(data)

	modityFunc(data, t)

	t.Log(data)

}

func modityFunc(arr [3]int, t *testing.T) {
	arr[2] = 20
	t.Logf("Arr in func=%v,arr addr", arr)
}
