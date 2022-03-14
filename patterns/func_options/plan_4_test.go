package func_options

import "testing"

type FuncOption func(e *MatchEngine)

func WithLogger() FuncOption {
	return func(e *MatchEngine) {
		e.Logger = &FileLogger{}
	}
}

func NewEngineWithOptions(driverId int, orderIds []int, options ...FuncOption) MatchEngine {
	engine := MatchEngine{DriverId: driverId, OrderIds: orderIds}

	for _, option := range options {
		option(&engine)
	}
	return engine
}

func TestCreateMatchEngineWithFuncOption(t *testing.T) {
	//直接通过开放属性来设置
	engine := NewEngineWithOptions(100, []int{199, 200}, WithLogger())
	engine.Run()

	//缺点:
	//1.需要学习成本

	//更进一步WithLogger可以变成带参数的
	//func WithLogger(logger) FuncOption {
	//	return func(e *MatchEngine) {
	//		e.Logger = logger
	//	}
	//}
}
