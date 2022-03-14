package func_options

import (
	"fmt"
)

type logger interface {
	log()
}

type FileLogger struct {
}

func (f FileLogger) log() {
	//write the disk file
}

type DBLogger struct {
}

func (D DBLogger) log() {
	//write to db
}

type matchConfig struct {
}

type MatchEngine struct {
	DriverId        int
	OrderIds        []int
	Rules           []string     //匹配规则
	Timeout         int          //匹配超时时间
	MaxGoroutineNum int          //最大使用协程数量
	Logger          logger       //记录日志
	Config          *matchConfig //匹配配置
}

func (e MatchEngine) Run() error {
	fmt.Println("match engine running")
	return nil
}

//plan2
func NewMatchEngine(driverId int, orderIds []int) MatchEngine {
	return MatchEngine{DriverId: driverId, OrderIds: orderIds}
}

func NewMatchEngineWithRules(driverId int, orderIds []int, rules []string) MatchEngine {
	return MatchEngine{DriverId: driverId, OrderIds: orderIds, Rules: rules}
}
