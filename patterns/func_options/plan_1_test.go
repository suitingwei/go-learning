package func_options

import "testing"

func TestCreateMatchEngine(t *testing.T) {
	//直接通过开放属性来设置
	engine := MatchEngine{
		DriverId:        0,
		OrderIds:        nil,
		Rules:           nil,
		Timeout:         0,
		MaxGoroutineNum: 0,
		Logger:          nil,
		Config:          nil,
	}
	engine.Run()

	//缺点:
	//1.没有统一的入口，重构的时候比较困难,比如新增一个shouldLog bool ，如果忘记赋值，那么很可能一直是默认的false,但是程序可以运行
	//2.没有构造函数，当需要设置某些默认值的时候，无法处理,比如默认logger，默认config
	//3.没有构造函数，无法在入口进行参数合法性校验，比如max_go_routine_num等
	//4.调用方心智负担大，不确定那些是必传参数
}
