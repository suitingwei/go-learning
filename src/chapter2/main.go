package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func learnHttp() {
	resp, err := http.Get("http://wordpress.alone-night.top")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	fmt.Printf("Response Status Code:%s\n", resp.StatusCode)

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))

	//var bodyBytes []byte

	//for {
	//	bytes,err :=  resp.Body.Read(bodyBytes)
	//
	//	if err!= nil || bytes == 0{
	//		fmt.Printf("Error occurred during the read operation of the response body:%s\n",err)
	//		break
	//	}
	//
	//	fmt.Printf("\tRead Bytes[%d]:%s\n",bytes,bodyBytes)
	//}

	return
}

func learnFile() {
	startTime := time.Now()
	file, err := os.OpenFile("./temp.file", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0755)

	defer file.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	sum := 0
	total := 1024 * 1024 * 1024
	for {
		bytes, err := file.WriteString(strings.Repeat("Same words every day", 96) + "\n")

		if err != nil {
			fmt.Printf("Writing files failed:%s\n", err)
			return
		}
		sum += bytes

		//降低的刷新频率，每500一次
		if sum%500 == 0 {
			showStatusbar(sum, total)
		}

		//1b * 1024 -> 1kb * 1024 ->  1mb * 1024 -> 1gb
		if sum >= total {
			endTime := time.Now()
			fmt.Printf("This program has written 1GB file within:%f seconds!", endTime.Sub(startTime).Seconds())
			return
		}
	}

}

func showStatusbar(currentValue int, total int) {
	//clear the current screen
	fmt.Println("\033[2J")

	//假如是：300/1000,那么要有30%是I,70%是.
	completePercent := 100 * currentValue / total
	unCompletePercent := 100 - completePercent
	fmt.Printf("Write [%d/%d] bytes: %s%s\n", currentValue, total, strings.Repeat("I", completePercent), strings.Repeat(".", unCompletePercent))
}

func main() {
	//learnFile()

}
