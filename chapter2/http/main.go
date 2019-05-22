package main

import (
	"fmt"
	"io/ioutil"
	"my-go-learnings"
	"net/http"
	"net/url"
	"reflect"
)

const BASE_DIDI_RECRUIT_URL = "http://talent.didiglobal.com/recruit-portal-service/api/job/front/list"

func main() {

	url := buildUrl(&my_go_learnings.DiDiRecruitRequestParams{
		JobType:     1,
		Page:        1,
		WorkArea:    "北京",
		Size:        10,
		RecruitType: 1,
	})
	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	resp, err := http.Get(urlPath)
	defer resp.Body.Close()
	s, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(s))

}

func buildUrl(didiParams *my_go_learnings.DiDiRecruitRequestParams) string {
	params := url.Values{}

	Url, err := url.Parse("http://baidu.com?fd=fdsf")
	if err != nil {
		panic(err.Error())
	}

}

func setHeaders(request *http.Request, bang *my_go_learnings.GeekBang) {
	ref := reflect.ValueOf(bang).Elem()

	for i := 0; i < ref.NumField(); i++ {
		valueField := ref.Field(i)
		typeField := ref.Type().Field(i)
		//fmt.Println(typeField.Name,valueField.Interface())
		request.Header.Add(typeField.Name, valueField.String())
	}
}
