//指针接受者 vs 普通接受者，以及struct的函数调用
//https://go.dev/doc/faq#pass_by_value
//https://go.dev/doc/faq#methods_on_values_or_pointers
package structs

import (
	"fmt"
	"testing"
)

type Person struct {
	cachedName *string
}

func (p *Person) GetName1() string {
	return p.getNameWithCache()
}

func (p *Person) GetName2() string {
	return p.getNameWithCache()
}

func (p *Person) getNameWithCache() string {
	if p.cachedName != nil {
		return *p.cachedName
	}
	fmt.Println("trigger time-consuming RPC")
	result := "mock result"
	p.cachedName = &result
	return result
}

func TestGetNameWithCache(t *testing.T) {
	p := Person{}

	t.Log(p.GetName1())
	t.Log(p.GetName2())
}

func TestNewWithMake(t *testing.T) {
	p := new(Person)
	arr := new([]Person)
	*arr = append(*arr, Person{})

	fmt.Println(p, p.GetName1())
	fmt.Printf("%+v\n", arr)
}
