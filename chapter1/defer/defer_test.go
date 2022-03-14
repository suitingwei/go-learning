package _defer

import (
	"fmt"
	"testing"
	"time"
)

func TestDeferParameters(t *testing.T) {
	//尽管defer会在函数结束的时候才会触发，但是他的传入参数早就做好了镜像了
	defer func(t time.Time) {
		fmt.Printf("The defer function triggered at :%s\n", t.Format("2006-01-02 15:04:05"))
	}(time.Now())

	//让这个函数在 5 秒后才结束
	time.Sleep(time.Second * 3)

	fmt.Printf("The function ended at:%s\n", time.Now().Format("2006-01-02 15:04:05"))

	output := []byte{1, 2, 3}

	newOutput := string(output[:])

	fmt.Println(newOutput)
}

//同一个函数内部，FILO
//这是有实际含义的，比如说打开一个文件描述符，在最上面defer close
//可能后续的流程里还会用到，那么肯定要最后执行这个defer
func TestMultipleDefersInSameFunc(t *testing.T) {
	defer func() {
		t.Log("first defer call")
	}()

	defer func() {
		t.Log("second defer call")
	}()

	t.Log("Normally response")

	//output
	//normally response
	//second defer call
	//first defer call
}

func TestDeferInNestedFuncCall(t *testing.T) {
	defer func() {
		t.Log("defer in the main func")
	}()

	a(t)
}

func a(t *testing.T) {
	defer func() {
		t.Log("first defer call in func a")
	}()

	defer func() {
		t.Log("second defer call in func a")
	}()

	b(t)

	t.Log("Normally response call A")
}

func b(t *testing.T) {
	defer func() {
		t.Log("first defer call in func B")
	}()

	defer func() {
		t.Log("second defer call in func B")
	}()

	t.Log("Normally response call B")
}
