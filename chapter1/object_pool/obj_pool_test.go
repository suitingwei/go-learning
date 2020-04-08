package object_pool

import (
	"sync"
	"testing"
)

type Pool struct {
	pool  chan int
	mutex sync.Mutex
}

func TestObjPool(t *testing.T) {

}

//Sync没法做真正的对象池,因为sync.pool会被GC回收
//sync.Pool会有两个对象池。第一个是协程私有的对象池
//这个对象池其实只能放一个对象，如果协程的私有对象不存在
//的话，回去共享对象池获取对象，如果还不存在，那么会使用
//New函数进行创建
func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			t.Log("Creating new object")
			return 100
		},
	}

	data := pool.Get().(int)
	t.Logf("Get data from pool:%d\n", data)

	data2 := pool.Get().(int)
	t.Logf("Get data2 from pool:%d\n", data2)

	pool.Put(999)
	data3 := pool.Get().(int)
	t.Logf("Get data3 from pool:%d\n", data3)

	data3 = pool.Get().(int)
	t.Logf("Get data4 from pool:%d\n", data3)

}
