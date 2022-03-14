package main

import (
	"crypto/md5"
	"encoding/json"
	"log"
	"math"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"time"
)

//分析调用最多的函数
func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	//data := []string{}
	i := 0
	for {
		//data = append(data, fmt.Sprintf("hello,world,this is the i=%d time call the append func", i))
		i++
		//
		//time.Sleep(time.Second * 1)

		rand.Seed(time.Now().Unix())
		random := rand.Intn(5)

		if random >= 1 && random <= 2 {
			A()
		} else if random <= 3 {
			B()
		} else {
			C()
		}
	}
}

func A() {
	timeConsumeFunc(88888)
}
func B() {
	timeConsumeFunc(100000)
}
func C() {
	timeConsumeFunc(5000)
}

func timeConsumeFunc(times int) {

	type valueStruct struct {
		value float64
	}
	result := 0.0
	for i := 0.0; i < float64(times); i++ {
		result = math.Sqrt(result * i)
		val := valueStruct{value: result}
		data, _ := json.Marshal(val)
		md5.Sum(data)
	}
}
