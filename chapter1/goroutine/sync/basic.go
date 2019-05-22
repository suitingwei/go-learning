package sync

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type data struct {
	sync.Mutex
}

//如果把 Mutex 包裹在struct,那么这个 struct 的方法必须用指针接受者才能正常工作
//比如下面这个，如果换成func(d data)test(str string){} 那么就没锁上
func (d *data) test(str string) {
	d.Lock()

	defer d.Unlock()

	for i := 0; i < 5; i++ {
		fmt.Println(str + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func LearnSync() {
	var wg sync.WaitGroup
	wg.Add(2)

	var d data

	go func() {
		defer wg.Done()
		d.test("read")
	}()

	go func() {
		defer wg.Done()
		d.test("write")
	}()

	wg.Wait()
}
