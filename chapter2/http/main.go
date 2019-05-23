package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"sync"
)

const BASE_DIDI_RECRUIT_URL = "http://talent.didiglobal.com/recruit-portal-service/api/job/front/list"

const (
	BaseInfo string = "https://api.github.com/users/suitingwei"
	//添加用户、密码，防止 github 限流
	SearchUserApi string = "https://suitingwei:suitingwei123@api.github.com/search/users"
)

func main() {

	downloader := New()

	//Set the search name of the github api
	downloader.searchName = "sui"

	//init the users container
	users := downloader.SearchUsers()

	for page = 2; page <= totalPages; page++ {
		fmt.Printf("Current user count is:%d, total user count:%d,totalPages:%d\n", users.CurrentUsersCount(), users.TotalCount, totalPages)

		newUsers := downloader.SearchUsers()

		downloader.DownloadThroughChannel(newUsers)
	}

}

func learnBasicHttpGet() {
	resp, err := http.Get(BaseInfo)

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	var userInfo GithubUserInfo

	err = json.NewDecoder(resp.Body).Decode(&userInfo)

	fmt.Println(userInfo)

}

type T struct {
	s string
}

func learnLoopRange() {
	ts := []T{
		T{"1"}, T{"2"}, T{"3"},
	}

	wg := &sync.WaitGroup{}
	wg.Add(3)

	for _, t := range ts {
		//t 是一个引用（指针），指向了数组中某一个元素
		fmt.Printf("&t pointer %T %p %v\n", &t, &t, t)

		//当把t的指针传入之后，其实是传递了 t 这个变量容器的地址。
		go func(pt *T) {
			//这时候在打印的的话。首先由一个新的指针容器，他存放的是t的地址。指向了t，而 t 最后指向了数组最后一个元素
			fmt.Printf("&pt pointer %T %p, pt pointer %T %p %v\n", &pt, &pt, pt, pt, pt)
			wg.Done()
		}(&t)
	}

	wg.Wait()
}
