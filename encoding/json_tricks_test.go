package encoding

import (
	"encoding/json"
	"github.com/gookit/goutil/dump"
	"github.com/pkg/errors"
	"strconv"
	"testing"
)

//json编码技巧

type User struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Order struct {
	Oid     string `json:"oid"`
	Address string `json:"address"`
}
type OrderStat struct {
	SeenDrivers int `json:"seen_drivers"`
	Star        int `json:"star"`
}

var user1 = User{
	Email:    "12345@gmail.com",
	Name:     "Tom",
	Password: "WeatherIsGood",
}

//1.临时忽略字段,本次为忽略password
func TestTempOmit(t *testing.T) {
	data, _ := json.Marshal(struct {
		*User
		Password bool `json:"password,omitempty"`
	}{
		User: &user1,
	})

	dump.P(string(data))
}

//2.临时增加字段
func TestTempAddField(t *testing.T) {
	data, _ := json.Marshal(struct {
		*User
		Token    string `json:"token"`
		Password bool   `json:"password,omitempty"`
	}{
		User:  &user1,
		Token: "TOKEN_TEMPLATE",
	})

	dump.P(string(data))
}

//3.临时合并多个结构体
func TestMergeMultipleStructs(t *testing.T) {
	data, _ := json.Marshal(struct {
		*User
		*Order
		*OrderStat
	}{
		&user1,
		&Order{
			Oid:     "TVRRME1UVXdPRGczTmpjME5ERTBNak0xT0E9PQ==",
			Address: "China Town, Australia",
		},
		&OrderStat{
			SeenDrivers: 100,
			Star:        5,
		},
	},
	)
	dump.P(string(data))
}

//4.完整的json拆分多个结构体
func TestSplitMultipleStructs(t *testing.T) {
	var jsonData = `{
  "oid": "attila@attilaolah.eu",
  "address": "Attila's Blog",
  "seen_drivers": 6,
  "star": 1
}`
	order := &Order{}
	orderStat := &OrderStat{}
	json.Unmarshal([]byte(jsonData), &struct {
		*Order
		*OrderStat
	}{order, orderStat})

	dump.P(order, orderStat)
}

//5.兼容字符串数字
func TestStringNumeric(t *testing.T) {
	//注意使用",string"tag，只能支持从string转过来，如果入参还可能是纯数字，就会报错
	var jsonData = `{ "field": "100"}`

	//错误的json数据，这次的field又是数字了
	//var jsonData = `{"field" :100}`

	type Person struct {
		Field int `json:"field,string"`
	}

	person := &Person{}

	err := json.Unmarshal([]byte(jsonData), person)
	if err != nil {
		t.Fatalf("failed to unmarshal,err=%v", err)
		return
	}
	dump.P(person)
}

type ResponseUser struct {
	//额外包一层，避免直接使用 User实现 UnmarshalJSON的时候导致无限递归
	inner *User
}

func (u *ResponseUser) UnmarshalJSON(b []byte) error {
	//这里的入参b，代表的是json字符串里对应字段，冒号后面的所有数据，也就是自动带有双引号
	unquoted, err := strconv.Unquote(string(b))
	if err != nil {
		return errors.Wrap(err, "failed to unquoted")
	}

	inner := &User{}
	err = json.Unmarshal([]byte(unquoted), inner)
	if err != nil {
		return err
	}
	*u = ResponseUser{inner: inner}
	return nil
}

//6.解析嵌套的json字符串
func TestDecodeNestedJsonString(t *testing.T) {
	var jsonData = `{
    "errno":0,
    "errmsg":"success",
    "data":{
        "user":"{\"email\":\"xxx@didiglobal.com\",\"name\":\"xxx\",\"password\":\"you shall not see this\"}"
    }
}`

	type Response struct {
		Errno  int    `json:"errno"`
		ErrMsg string `json:"errmsg"`
		Data   struct {
			User ResponseUser `json:"user"`
		} `json:"data"`
	}

	response := &Response{}

	err := json.Unmarshal([]byte(jsonData), response)
	if err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
		return
	}
	dump.P(response)
}

//反转义json字符串
func TestUnQuote(t *testing.T) {
	var input = `{\"name\":\"suitingwei\",\"age\":13}`

	unquote, err := strconv.Unquote(`"` + input + `"`)
	if err != nil {
		t.Fatalf("failed to unquote: %v", err)
		return
	}
	dump.P(unquote)

	var input2 = `"{\"email\":\"xxx@didiglobal.com\",\"name\":\"xxx\",\"password\":\"you shall not see this\"}"`

	unquote2, err := strconv.Unquote(input2)
	if err != nil {
		t.Fatalf("failed to unquote: %v", err)
		return
	}
	dump.P(unquote2)

}
