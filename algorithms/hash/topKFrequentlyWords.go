package hash

import (
	"fmt"
	"sort"
)

type MapItem struct {
	word  string
	count int
}

type ValueSortMap struct {
	ResultData   []MapItem
	OriginalData []string
}

func (m *ValueSortMap) SortByValue() {
	temp := make(map[string]int)

	for _, value := range m.OriginalData {
		_, ok := temp[value]

		if ok {
			temp[value]++
		} else {
			temp[value] = 1
		}
	}

	for word, count := range temp {
		mapItem := MapItem{
			word:  word,
			count: count,
		}
		m.ResultData = append(m.ResultData, mapItem)
	}

	sort.SliceStable(m.ResultData, func(i, j int) bool {
		return m.ResultData[i].count >= m.ResultData[j].count
	})
}

func Solve() {

	testData := []string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}

	result := &ValueSortMap{
		OriginalData: testData,
	}

	result.SortByValue()

	for _, value := range result.ResultData {
		fmt.Printf("%s-%d\n", value.word, value.count)
	}

}
