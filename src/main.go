package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

var wg sync.WaitGroup

func writeFile(fileName string, times int) {

	startTime := time.Now()

	defer wg.Done()

	file, fileErr := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)

	if fileErr != nil {
		log.Fatalln("Failed to open file" + fileErr.Error())
	}

	defer func() {
		file.Close()

		endTime := time.Now()

		fmt.Printf("Time elapsed:%f(s)\n", endTime.Sub(startTime).Seconds())
	}()

	for i := 1; i < times; i++ {
		words := fmt.Sprintf("hello world:%d\n", i)

		_, err := file.WriteString(words)

		if err != nil {
			log.Println("Error occurred during file writing process:" + err.Error())
		}
	}

}

type Frequency map[string]int

func main() {
	//testData:=[]string{"bella","label","roller"}
	//testData := []string{"cool", "lock", "cook"}
	//
	//result := commonChars(testData)
	//
	//fmt.Printf("Common characters are:%v\n", result)

	testData := []int{18, 12, -18, 18, -19, -1, 10, 10}

	result := canThreePartsEqualSum(testData)
	fmt.Printf("Result is :%t\n", result)
}

func commonChars(A []string) []string {
	//计算每一个单词的字符出现频率
	allFrequencies := make([]Frequency, 0)

	for _, str := range A {
		allFrequencies = append(allFrequencies, charFre(str))
	}

	firstFre := allFrequencies[0]

	otherFrequencies := allFrequencies[1:]

	result := []string{}
	for char, fre := range firstFre {

		for _, frequencyMap := range otherFrequencies {
			newFrequency, ok := frequencyMap[char]
			if !ok {
				fre = 0
			} else if newFrequency < fre {
				fre = newFrequency
			}
		}

		for i := 0; i < fre; i++ {
			result = append(result, char)
		}
	}

	return result
}

func charFre(str string) Frequency {
	fre := Frequency{}

	for _, charInt := range str {
		char := string(charInt)
		_, ok := fre[char]

		if ok {
			fre[char]++
		} else {
			fre[char] = 1
		}
	}
	return fre
}

func canThreePartsEqualSum(nums []int) bool {

	totalSum := sumArray(nums)

	for i := 1; i < len(nums)-1; i++ {
		for j := len(nums) - 1; j >= i+1; j-- {

			sumA := sumArray(nums[0:i])
			sumB := sumArray(nums[i:j])
			sumC := sumArray(nums[j:])

			fmt.Printf("Temp result,total sum=%d:\n"+
				"\tnumsA:=%v,sum=%d\n\tnumsB:=%v,sum=%d\n\tnumsC:=%v,sum=%d\n",
				totalSum, nums[0:i], sumA, nums[i:j], sumB, nums[j:], sumC,
			)

			//终止条件
			if sumA == sumB && sumB == sumC {
				fmt.Printf("Found result:\n"+
					"\tnumsA:=%v\n\tnumsB:=%v\n\tnumsC:=%v\n",
					nums[0:i], nums[i:j], nums[j:],
				)
				return true
			}
		}
	}
	return false
}

func sumArray(nums []int) int {

	total := 0

	for _, value := range nums {
		total += value
	}

	return total
}
