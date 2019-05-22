package main

import (
	"encoding/json"
	"fmt"
	"log"
	"my-go-learnings/chapter2/http/models"
	"net/http"
)

const BASE_DIDI_RECRUIT_URL = "http://talent.didiglobal.com/recruit-portal-service/api/job/front/list"

const (
	BaseInfo      string = "https://api.github.com/users/suitingwei"
	SearchUserApi string = "https://api.github.com/search/users"
)

type GithubUserInfo struct {
	Login string `json:"login"`
	Id    int    `json:"id"`
}

type GithubUserList struct {
	Users            []GithubUserInfo `json:"items"`
	TotalCount       int              `json:"total_count"`
	InCompleteResult bool             `json:"incomplete_results"`
}

func main() {
	client := &http.Client{} //创建一个请求

	req, err := http.NewRequest(http.MethodGet, SearchUserApi, nil)

	if err != nil {
		log.Fatalln(err)
	}

	//创建这个请求的 query
	query := req.URL.Query()
	query.Add("q", "suitin")
	query.Add("sort", "joined")

	req.URL.RawQuery = query.Encode()

	fmt.Println(req.URL.String())

	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	var users GithubUserList
	err = json.NewDecoder(resp.Body).Decode(&users)

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(users)
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

func buildUrl(didiParams *models.DiDiRecruitRequestParams) string {
	return ""
}
