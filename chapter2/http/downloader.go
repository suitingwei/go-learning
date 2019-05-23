package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"
)

type GithubUserInfo struct {
	Login  string `json:"login"`
	Id     int    `json:"id"`
	Avatar string `json:"avatar_url"`
}

type GithubUserList struct {
	Users            []GithubUserInfo `json:"items"`
	TotalCount       int              `json:"total_count"`
	InCompleteResult bool             `json:"incomplete_results"`
}

//Current users count will return the current users in the struct.
//As a memory cache for the download procure.
func (users *GithubUserList) CurrentUsersCount() int {
	return len(users.Users)
}

//AppendUsers will append new users in the GithubUserList struct into the
//current struct.
func (users *GithubUserList) AppendUsers(newUsers *GithubUserList) {
	for _, newUser := range newUsers.Users {
		users.Users = append(users.Users, newUser)
	}
}

type Downloader struct {
}

func New() *Downloader {
	return &Downloader{}
}

//download method will download all the users' avatars in the given users slice.
//Be noticed about that the github api will return all user's count, be each page
//only return part of the users.
func (downloader Downloader) Download(users *GithubUserList) {
	//计算耗时
	startTime := time.Now()
	fmt.Printf("Start download %d avatars!\n", users.TotalCount)

	//使用协程并发下载哦
	wg := sync.WaitGroup{}
	wg.Add(users.TotalCount)

	for _, user := range users.Users {
		fmt.Println(user.Login, user.Avatar)
		go DownloadUrl(user, &wg)
	}
	wg.Wait()

	//输出耗时
	endTime := time.Now()
	fmt.Printf("Download %d avaters using %f seconds\n", users.TotalCount, endTime.Sub(startTime).Seconds())
}

//Request method will request the github api to search users.
func (downloader Downloader) SearchUsers(name string, page int) *GithubUserList {
	client := &http.Client{} //创建一个请求

	req, err := http.NewRequest(http.MethodGet, SearchUserApi, nil)

	if err != nil {
		log.Fatalln(err)
	}

	//创建这个请求的 query
	query := req.URL.Query()
	query.Add("q", name)
	query.Add("page", strconv.Itoa(page))
	query.Add("sort", "joined")

	req.URL.RawQuery = query.Encode()

	fmt.Printf("Github api url: %s\n", req.URL.String())

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

	return &users
}

//DownloadAvatar method will download the users' avatars through goroutines.
func DownloadUrl(user GithubUserInfo, wg *sync.WaitGroup) {
	defer wg.Done()
	imagePath, err := filepath.Abs("./")
	if err != nil {
		log.Fatalf("Failed to obtain the image path:%s\n", err.Error())
	}

	imagePath = filepath.Join(imagePath, "githubUserAvatars")

	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		_ = os.Mkdir(imagePath, os.ModePerm)
	}

	file, err := os.Create(filepath.Join(imagePath, user.Login+".jpg"))

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Start to downloading...")
	resp, err := http.Get(user.Avatar)

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	_, err = io.Copy(file, resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("User:%s's avatar:[%s] has been downloaded successfully!\n", user.Login, user.Avatar)
}
