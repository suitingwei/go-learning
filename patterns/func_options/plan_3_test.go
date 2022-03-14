package func_options

import "testing"

type Option struct {
	Logger logger
	Config *matchConfig
}

func NewEngineWithOption(driverId int, orderIds []int, option Option) MatchEngine {
	engine := MatchEngine{DriverId: driverId, OrderIds: orderIds}
	if option.Logger != nil {
		engine.Logger = option.Logger
	}

	if option.Config != nil {
		engine.Config = option.Config
	}
	return engine
}

func TestCreateMatchEngineWithOption(t *testing.T) {
	//直接通过开放属性来设置
	engine := NewEngineWithOption(100, []int{1288, 0100}, Option{
		Logger: nil,
		Config: nil,
	})
	engine.Run()

	//缺点:
	//1.Option的默认值和零值有可能会分不清
}
