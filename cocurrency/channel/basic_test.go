package channel

import "testing"

func TestName(t *testing.T) {
	dataLength := 20
	result := make(chan int, dataLength)

	for i := 0; i < dataLength; i++ {

		go func(num int) {
			//故意丢失一些写操作
			if num%3 == 0 {
				result <- num * 10
			}
		}(i)
	}

	for i := 0; i < dataLength; i++ {
		t.Log(<-result)
	}
	//output : fatal error: all goroutines are asleep - deadlock!

	//for resultNum := range result {
	//	t.Log(resultNum)
	//}
	//fatal error: all goroutines are asleep - deadlock!

}

func TestReadWriteSameTimes(t *testing.T) {
	dataLength := 20
	result := make(chan int, dataLength)

	for i := 0; i < dataLength; i++ {

		go func(num int) {
			//故意丢失一些写操作
			if num%3 == 0 {
				result <- num * 10
			} else {
				result <- num
			}
		}(i)
	}

	for i := 0; i < dataLength; i++ {
		t.Log(<-result)
	}

	//for resultNum := range result {
	//	t.Log(resultNum)
	//}
	//fatal error: all goroutines are asleep - deadlock!

}
