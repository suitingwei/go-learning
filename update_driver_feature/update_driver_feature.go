package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
)

const Url = "https://gsapi-sim118-us01-test.intra.xiaojukeji.com/gulfstream/driver/v2/driver/dSetTickPickerMode"

const hallQuickSwitchOn = "1"
const hallQuickSwitchOff = "0"

func main() {
	var wg sync.WaitGroup
	fileList := []string{
		"./Cuernavaca.csv",
		"./Mazatlan.csv",
	}
	for _, file := range fileList {
		wg.Add(1)
		go updateCity(file, hallQuickSwitchOn, &wg)
	}
	wg.Wait()
}

func updateCity(filePath string, hallQuickSwitch string, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()

	csvFile, err := os.OpenFile(filePath, os.O_RDONLY, 0666)

	if err != nil {
		panic(fmt.Sprintf("failed to open file=%s,err=%s", filePath, err))
	}

	csvReader := csv.NewReader(csvFile)

	line := 1

	for {
		fields, err := csvReader.Read()
		if err == io.EOF {
			fmt.Printf("file=%s has finished", filePath)
			return
		}

		if err != nil {
			panic(fmt.Sprintf("failed to read csv row file=%s,err=%s", filePath, err))
		}

		//fields[0]是司机id
		updateDriver(fields[0], hallQuickSwitch, line)
		line++
	}

}

func updateDriver(did string, hallQuickSwitch string, line int) {
	//fmt.Printf("updating driver feature,driverId=%s,hall_quick_switch=%s\n", did, hallQuickSwitch)
	fmt.Printf("currentProcessingLine=\r%d", line)
	payload := strings.NewReader(`{"common":{"ticket":"YR1NNZRJDp_rpk9cwCVrHwj2XVhId9Ri5HkcXvmVXBEszDsOwjAQBuG7TMsq2t-OHXlLjsANeIRHYyQQVZS7IxSqmepb6E6QBx8co4uQ0RMxueu3mXCjj1sKwf6AcSTAOBG1eJM3jXmqas24ELkaM7Hwfn5e5_mvrcaVUE2qYy5Jxo1gV0pSayotKwnjvtEPwtdvAAAA__8=","terminal_id":5,"app_version":"7.6.13","origin_id":5,"lat":-22.7622926,"lng":-42.853846,"product_id":316,"biz_type":2,"utc_offset":-180,"lang":"pt-BR","location_cityid":55000161,"location_country":"BR","trip_type":"","map_type":"wgs84","data_type":"","mobile_type":"","network_type":"LTE","client_type":0,"channel":0,"platform_type":0,"a3_token":"uic\/k3kkcMYr7LzH9ogk\/kw8TNt16VUTVsWcFeeFUoCPNro1TjcEjFJrZ3fwjIUkdwY5rwp5TMgwJaILZGTPLUHrJ1tZTfcgeMzzZnijfJk8NHahLyYj0zrgI7cjBDWLbgqhBlqGE7w=","wsgenv":"eV60A6lSCIlLXhhhAAAAA6ABAAA85XwVjgEe8btHN5jGbHzHd+4sy0d6PjQZIYbtmiOCu2ijFvxvNtDTa9NxSxRctLHLYSZPy89h2FqiNc99LgE15DkFvW5p6VqY4fWzMcSZjl67Vg4fm5FnuR\/k52yldV1VkgBIbn7Fd7fh4LdzrVYzkKihlxWfPkla4t42lcgqUK7\/i9Vj7MqvREpbUmlT9cV+9kKrTynEB2gQK+41CN2LkPgTwFB4nMOdgAMPA4p+Ki0kVE7A94IeA27OgZoHRpLsxewOjDr4ZLhliPKnLZ9NNxxnjYanqj5MynkpbrhNFIWD8rIl8TjtFmftWQuUo27TRoBL41YjiWV8a97uQ00J+wgdK3orAHEsa+uvrx21ddUUzmg\/nLCfPWn30xCvFZEm8eSwUwwyAAcDvKo\/KL7FSdHx1sS+GIT5EhXq1ZOYLACocpo5ZhPE\/5Jit2kmzCT0a5t44hhBcWYtAWkX4rYkvKwFf+EvGfHPFmuRkJ4h7VoVj0Y8E0j0anEaBDfeyIjH1YQlsSU4ESAGvoaMOBRRrL7Tbv6l6VreKZ1zsOwSiS7QXvHuLZNGqCxzo4yS6HFa7RJuLGNGRLh58BBD6EUZmUgZnmFxKA0Y7Oow1EdZz\/","device_info":{"uuid":"","suuid":"","idfa":"","imei":"","imsi":"","pixels":"","deviceid":"2324e87e48202c728bebd7480219c30f","iccid":"","sdid":"","model":"moto g(9) play","brand":"","os":"11","osType":"android","osVersion":""}},"login":{"uid":` + did + `,"role":2,"pid":0,"driver_id":` + did + `,"phone":"+5521991593121","country_code":"+55","is_test_cell":false},"config":{"signpost":null},"request":{"caller":"null","hall_quick_switch":` + hallQuickSwitch + `}}`)

	client := &http.Client{}
	req, err := http.NewRequest("POST", Url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Host", "gsapi-sim118-us01-test.intra.xiaojukeji.com")
	req.Header.Add("tripcountry", "UG")
	req.Header.Add("didi-header-hint-content", "{\"trip_country\":\"UG\",\"locale\":\"en_UG\",\"lang\":\"en-US\",\"location_cityid\":256999900}")
	req.Header.Add("user-agent", "Android/9 didihttp OneNet/2.1.1.10.01 com.canoe.driver/1.0.0")
	req.Header.Add("cityid", "256999900")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("didi-header-rid", "ac189c2b610a34ce000051921fenqw2222333")
	req.Header.Add("wsgsig", "dd04-tj93Ha6E/JxUhF3v2PB7sDGG+vWTGT1/FrHpdmNsxObfFpfyXb4vfklz/TLzP0XjG3xzWtWRjuMhiR9H+3bqBO0PTTMtu4HKN7spoDQO/TTDQO0JaJIxVd7siawpRXU4auoFDdKJSAl0QRd8auyoVnbpiMVGW1gebMo0XkH")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))

}
