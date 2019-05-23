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
	"sync/atomic"
	"time"
)

type Downloader struct {
	searchName      string              // the search name of the github api
	currentPage     int                 // the current request page
	mainChannel     chan GithubUserInfo // main channel receive from the main routine
	downloadedCount int64               // the total downloaded avatars
	wg              sync.WaitGroup      // the sync group
	totalCount      int                 // the total avatars to be downloaded
	PageSize        int                 // the page size of the github search user api
}

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

func New() *Downloader {
	downloader := &Downloader{
		mainChannel:     make(chan GithubUserInfo),
		downloadedCount: 1,
		wg:              sync.WaitGroup{},
	}

	go func() {
		logFile, _ := os.Create("./mainRoutine.log")
		log.SetOutput(logFile)
		log.SetFlags(log.LstdFlags | log.Ltime)

		for {
			log.Println("waiting for the main channel's data coming....")
			user := <-downloader.mainChannel
			go downloader.DownloadUrl(user)
		}
	}()

	return downloader
}

//download method will download all the users' avatars in the given users slice.
//Be noticed about that the github api will return all user's count, be each page
//only return part of the users.
func (downloader *Downloader) Download(users *GithubUserList) {
	//计算耗时
	startTime := time.Now()
	fmt.Printf("Start download %d avatars!\n", users.TotalCount)

	//使用协程并发下载哦
	wg := sync.WaitGroup{}
	wg.Add(users.TotalCount)

	for _, user := range users.Users {
		go downloader.DownloadUrl(user)
	}
	wg.Wait()

	//输出耗时
	endTime := time.Now()
	fmt.Printf("Download %d avaters using %f seconds\n", users.TotalCount, endTime.Sub(startTime).Seconds())
}

//Request method will request the github api to search users.
func (downloader *Downloader) SearchUsers() *GithubUserList {
	client := &http.Client{} //创建一个请求

	req, err := http.NewRequest(http.MethodGet, SearchUserApi, nil)

	if err != nil {
		log.Fatalln(err)
	}

	//创建这个请求的 query
	query := req.URL.Query()
	query.Add("q", downloader.searchName)
	query.Add("page", strconv.Itoa(downloader.currentPage))
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

	//This page's request is successfully, go to the next page
	downloader.currentPage++

	if downloader.PageSize != 0 {
		//update the downloader's total count attr.
		downloader.PageSize = users.CurrentUsersCount()

		//set the total count of the search users.
		downloader.totalCount = users.TotalCount

		//Add the wg
		downloader.wg.Add(users.TotalCount)
	}

	return &users
}

//Download through will pass the new user into the downloader's main channel
func (downloader *Downloader) DownloadThroughChannel(newUsers *GithubUserList) {
	for _, newUser := range newUsers.Users {
		downloader.mainChannel <- newUser
	}
}

//DownloadAvatar method will download the users' avatars through goroutines.
func (downloader *Downloader) DownloadUrl(user GithubUserInfo) {
	defer downloader.wg.Done()
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
		log.Println(err)
		return
	}
	resp, err := http.Get(user.Avatar)

	if err != nil {
		log.Println(err)
		return
	}

	defer resp.Body.Close()

	_, err = io.Copy(file, resp.Body)

	if err != nil {
		log.Println(err)
		return
	}

	//atomic.AddInt64(downloadedCount,1)
	atomic.AddInt64(&downloader.downloadedCount, 1)

	fmt.Printf("[%d/%d]\t[%s] has been downloaded successfully!\n", atomic.LoadInt64(&downloader.downloadedCount), downloader.totalCount, user.Avatar)
}
