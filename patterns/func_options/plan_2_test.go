package func_options

import "testing"

func TestCreateMatchEngineWithCtr(t *testing.T) {
	//直接通过开放属性来设置
	engine := NewMatchEngine(100, []int{1100, 200})
	engine.Run()

	rulesEngine := NewMatchEngineWithRules(100, []int{1100, 200}, []string{"anti-spam", "eda-rules"})
	rulesEngine.Run()

	//缺点:
	//1.如果在一个大的构造函数里，那么随着功能增多，这个构造函数会越来越长
	//2.因为golang不支持可选参数，当某一个参数不是必传，那么会导致一堆默认值
}
