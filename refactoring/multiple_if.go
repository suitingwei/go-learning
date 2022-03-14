package refactoring

import "fmt"

type OrderInfo struct {
	ProductId   int32
	ExtraType   int
	OrderStatus int32
}

const (
	jpProductId      = 100
	fastCarProductId = 200
)

const (
	carpoolExtraType  = 1
	openRideExtraType = 2
)

func codeSmell(o OrderInfo) error {
	if isFastCar(o.ProductId) {
		if isCarpool(o.ExtraType) {
			return business_func1(o)
		} else if isOpenRide(o.ExtraType) {
			return business_func2(o)
		} else {
			return default_business_func(o)
		}
	}

	if isJapanese(o.ProductId) {
		if isCarpool(o.ExtraType) {
			return jp_carpool_func(o)
		} else if isOpenRide(o.ExtraType) {
			return business_jp_func2(o)
		} else {
			return default_business_jp_func(o)
		}
	}
	return nil
}

type Conditions struct {
	IsJapanese bool
	IsFastCar  bool
	IsCarpool  bool
	IsOpenRide bool
}

type handleFunc func(o OrderInfo) error

var conditionMapping = map[Conditions]handleFunc{
	{
		IsJapanese: true,
		IsFastCar:  false,
		IsCarpool:  true,
		IsOpenRide: false,
	}: jp_carpool_func,
	{
		IsJapanese: true,
		IsFastCar:  false,
		IsCarpool:  false,
		IsOpenRide: true,
	}: jp_openride_func,
	{
		IsJapanese: false,
		IsFastCar:  false,
		IsCarpool:  true,
		IsOpenRide: true,
	}: jp_carpool_func,
	{
		IsJapanese: false,
		IsFastCar:  false,
		IsCarpool:  true,
		IsOpenRide: true,
	}: jp_carpool_func,
}

func Optimize(o OrderInfo) error {
	cond := Conditions{
		IsJapanese: isJapanese(o.ProductId),
		IsFastCar:  isFastCar(o.ProductId),
		IsCarpool:  isCarpool(o.ExtraType),
		IsOpenRide: isOpenRide(o.ExtraType),
	}

	if handler, ok := conditionMapping[cond]; ok {
		return handler(o)
	}
	return default_business_func(o)
}

func default_business_jp_func(o OrderInfo) error {
	fmt.Println("running default business jp func")
	return nil
}

func business_jp_func2(o OrderInfo) error {
	fmt.Println("running business jp func_2")
	return nil
}

func jp_carpool_func(o OrderInfo) error {
	fmt.Println("日本拼车逻辑")
	return nil
}

func isJapanese(id int32) bool {
	return id == jpProductId
}

func default_business_func(o OrderInfo) error {
	fmt.Println("running default business func")
	return nil
}

func business_func2(o OrderInfo) error {
	fmt.Println("running business func_2")
	return nil
}

func isOpenRide(extraType int) bool {
	return extraType == openRideExtraType
}

func jp_openride_func(o OrderInfo) error {
	fmt.Println("日本open_ride逻辑")
	return nil
}

func business_func1(o OrderInfo) error {
	fmt.Println("running business func_1")
	return nil
}

func isFastCar(id int32) bool {
	return id == fastCarProductId
}
func isCarpool(extraType int) bool {
	return extraType == carpoolExtraType
}

var weekdayMap = map[int]string{
	1: "Monday",
	2: "Tuesday",
	3: "Wednesday",
	4: "Thursday",
	5: "Firday",
	6: "Saturday",
	7: "Sunday",
}

func getWeekName(d int) string {
	if result, ok := weekdayMap[d]; ok {
		return result
	}
	return "Error"
}
