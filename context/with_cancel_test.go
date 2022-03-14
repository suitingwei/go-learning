package context

import (
	"context"
	"fmt"
	"testing"
)

func TestWithCancel(t *testing.T) {
	cancelCtx, cancel := context.WithCancel(context.Background())

	for val := range getNumbers(cancelCtx) {
		//做好某个限制条件，达到标准就执行cancel
		if val > 5 {
			//这里注意cancel里的注释，cancel只是负责通知子协程，其他什么都不做
			cancel()
			//这里必须要break，否则cancel()不会跳出这个循环。 那么最终整个for循环，都会阻塞在 getNumbers() 从管道获取数据
			//但是也获取不到数据，最终main_go_routine会阻塞，但是child_go_routine因为收到了cancel_signal，所以也退出了
			//最终就死锁了
			break
		} else {
			t.Log(val)
		}
	}
}

func getNumbers(ctx context.Context) <-chan int {
	//标准生成器模式，返回一个result_channel，然后内部开启一个协程来处理其他任务，给result_channel传递数据
	resultChan := make(chan int)

	//生成器业务逻辑
	n := 0
	go func() {
		for {
			//必须用select多路查询，获取是否有done信号，如果直接if ctx.Done()就会阻塞当前的for循环，最终啥也不干只能等cancel
			select {
			case <-ctx.Done():
				{
					fmt.Println("received cancellation signal")
					return
				}
			default:
				//默认情况下，把生成的数据，写回channel，等到其他读协程拉取数据
				//这里注意，虽然是for循环无限循环，但是如果没有协程从channel读取数据，一样会彻底阻塞
				resultChan <- n
				n++
			}
		}
	}()

	//返回改协程
	return resultChan
}

func TestWithValue(t *testing.T) {
	ctx := context.Background()

	ctx = context.WithValue(ctx, "Name", "suitingwei")
	ctx = context.WithValue(ctx, "Age", "18")

	t.Logf("%v", ctx)
}
